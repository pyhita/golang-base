#!/usr/bin/env bash
#
# mount_propagation_test.sh
#
# 对比实验：验证 Linux mount 传播机制 private / slave / shared 三种模式，
# 对应 Kubernetes 里的 None / HostToContainer / Bidirectional。
#
# 思路：
#   1. 在"宿主机"上准备一个挂载点 HOST_DIR，设置传播类型
#   2. 用 unshare --mount 起一个长驻的后台进程，代表"容器"的 mount namespace，
#      在它里面把同一个挂载点也设置成对应的传播类型
#   3. 用 nsenter -t <容器PID> -m 分别在"容器视角"里执行挂载操作，
#      用普通命令在"宿主机视角"里执行挂载操作
#   4. 互相检查对方新增的挂载是否可见
#
# 需要 root 权限。仅建议在测试机 / 容器 / 虚拟机里跑，不要在生产宿主机上跑。
#
# 用法:
#   sudo bash mount_propagation_test.sh private
#   sudo bash mount_propagation_test.sh slave
#   sudo bash mount_propagation_test.sh shared
#   sudo bash mount_propagation_test.sh all
#

set -uo pipefail

if [[ $EUID -ne 0 ]]; then
    echo "错误：必须以 root 运行 (sudo bash $0 ...)" >&2
    exit 1
fi

for cmd in unshare nsenter findmnt mountpoint; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
        echo "错误：缺少命令 '$cmd'，请先安装 util-linux 相关工具" >&2
        exit 1
    fi
done

BASE=/tmp/mount-propagation-demo
HOST_DIR="$BASE/shared_root"
FROM_HOST_SUBDIR="from_host"           # 宿主机侧新增的挂载
FROM_CONTAINER_SUBDIR="from_container" # 容器侧新增的挂载

CONTAINER_PID=""

cleanup() {
    if [[ -n "$CONTAINER_PID" ]] && kill -0 "$CONTAINER_PID" 2>/dev/null; then
        kill "$CONTAINER_PID" 2>/dev/null || true
        wait "$CONTAINER_PID" 2>/dev/null || true
    fi
    umount "$HOST_DIR/$FROM_CONTAINER_SUBDIR" 2>/dev/null || true
    umount "$HOST_DIR/$FROM_HOST_SUBDIR" 2>/dev/null || true
    umount "$HOST_DIR" 2>/dev/null || true
    rm -rf "$BASE" 2>/dev/null || true
    CONTAINER_PID=""
}
trap cleanup EXIT

usage() {
    echo "用法: $0 [private|slave|shared|all]"
    exit 1
}

k8s_name() {
    case "$1" in
        private) echo "None" ;;
        slave)   echo "HostToContainer" ;;
        shared)  echo "Bidirectional" ;;
    esac
}

# 宿主机侧的传播类型：slave 必须依附一个 shared 的"主"挂载点才能生效，
# 所以 HostToContainer 场景下，宿主机侧本身要设为 shared，容器侧才设为 slave。
host_propagation_for() {
    case "$1" in
        private) echo "private" ;;
        slave)   echo "shared" ;;   # 宿主机是"主"，必须 shared
        shared)  echo "shared" ;;
    esac
}

MODE="${1:-}"
[[ -z "$MODE" ]] && usage

run_one_mode() {
    local mode="$1"
    echo
    echo "=============================================="
    echo "  测试模式: $mode  (K8s 对照: $(k8s_name "$mode"))"
    echo "=============================================="

    cleanup
    mkdir -p "$HOST_DIR"
    mkdir -p "$HOST_DIR/$FROM_HOST_SUBDIR" "$HOST_DIR/$FROM_CONTAINER_SUBDIR"

    local host_prop
    host_prop="$(host_propagation_for "$mode")"

    # 1. 让 HOST_DIR 成为一个挂载点（bind 自己），然后设置传播类型（宿主机侧）
    #    注意：宿主机侧必须是 shared 才能作为"主"向 slave 传播；private 场景下双方都是 private。
    mount --bind "$HOST_DIR" "$HOST_DIR"
    mount --make-"$host_prop" "$HOST_DIR"

    echo
    echo "[宿主机] $HOST_DIR 传播类型 (设置值: $host_prop):"
    findmnt -o TARGET,PROPAGATION "$HOST_DIR" | sed 's/^/    /'

    # 2. 起一个长驻后台进程，代表"容器"的独立 mount namespace
    #    注意：unshare --mount 默认会把新 namespace 里的挂载点强制标记为 private（安全默认行为），
    #    必须显式加 --propagation unchanged，才能保留宿主机侧原有的 shared 对等组关系，
    #    否则后面 mount --make-slave 会因为没有"主"而失效（回退成 private）。
    unshare --mount --propagation unchanged sleep 3600 &
    CONTAINER_PID=$!
    sleep 0.3   # 等 namespace 真正建立

    echo
    echo "[容器] 已创建 mount namespace, PID=$CONTAINER_PID"

    # 容器侧按目标 mode 设置（private/slave/shared 分别对应 None/HostToContainer/Bidirectional）
    nsenter --mount="/proc/$CONTAINER_PID/ns/mnt" -- mount --make-"$mode" "$HOST_DIR" 2>/dev/null || true

    echo "[容器] $HOST_DIR 传播类型 (容器视角):"
    nsenter --mount="/proc/$CONTAINER_PID/ns/mnt" -- findmnt -o TARGET,PROPAGATION "$HOST_DIR" | sed 's/^/    /'

    # 3. 在【宿主机】侧新增一个挂载
    echo
    echo "-- 在【宿主机】侧挂载 tmpfs 到 $FROM_HOST_SUBDIR --"
    mount -t tmpfs tmpfs "$HOST_DIR/$FROM_HOST_SUBDIR"

    # 检查【容器】侧是否可见 —— 对应 HostToContainer 方向
    if nsenter --mount="/proc/$CONTAINER_PID/ns/mnt" -- mountpoint -q "$HOST_DIR/$FROM_HOST_SUBDIR"; then
        echo "[容器] 结果: 能看到宿主机新增挂载  ==> 宿主机->容器 可见 (对应 HostToContainer 生效)"
    else
        echo "[容器] 结果: 看不到宿主机新增挂载  ==> 宿主机->容器 不可见"
    fi

    # 4. 在【容器】侧新增一个挂载
    echo
    echo "-- 在【容器】侧挂载 tmpfs 到 $FROM_CONTAINER_SUBDIR --"
    nsenter --mount="/proc/$CONTAINER_PID/ns/mnt" -- mount -t tmpfs tmpfs "$HOST_DIR/$FROM_CONTAINER_SUBDIR"

    # 检查【宿主机】侧是否可见 —— 对应 Bidirectional 方向
    if mountpoint -q "$HOST_DIR/$FROM_CONTAINER_SUBDIR"; then
        echo "[宿主机] 结果: 能看到容器新增挂载  ==> 容器->宿主机 可见 (对应 Bidirectional 生效)"
    else
        echo "[宿主机] 结果: 看不到容器新增挂载  ==> 容器->宿主机 不可见"
    fi

    echo
    echo "-- 宿主机视角下 $HOST_DIR 的完整挂载树 --"
    findmnt -R "$HOST_DIR" | sed 's/^/    /'

    cleanup
}

case "$MODE" in
    private|slave|shared)
        run_one_mode "$MODE"
        ;;
    all)
        run_one_mode private
        run_one_mode slave
        run_one_mode shared
        ;;
    *)
        usage
        ;;
esac

trap - EXIT
echo
echo "=============================================="
echo "  实验结束，对照关系："
echo "    private -> K8s None            : 双向都看不到（最安全，默认值）"
echo "    slave   -> K8s HostToContainer : 宿主机->容器 可见，容器->宿主机 不可见"
echo "    shared  -> K8s Bidirectional   : 双向都可见（需要特权容器，风险最高）"
echo "=============================================="