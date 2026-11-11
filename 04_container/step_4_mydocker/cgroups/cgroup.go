package cgroups

import (
	"github.com/pyhita/mydocker/cgroups/resource"
	log "github.com/sirupsen/logrus"
)

type CgroupManager interface {
	Apply(pid int) error
	Set(res *resource.ResourceConfig) error
	Destroy() error
}

func NewCgroupManager(path string) CgroupManager {
	if IsCgroup2UnifiedMode() {
		log.Infof("use cgroup v2")
		return NewCgroupV2Manager(path)
	}
	log.Infof("use cgroup v1")
	return NewCgroupV1Manager(path)
}
