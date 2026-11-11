package cgroups

import (
	"github.com/pyhita/mydocker/cgroups/fs2"
	"github.com/pyhita/mydocker/cgroups/resource"
	"github.com/sirupsen/logrus"
)

type CgroupV2Manager struct {
	Path       string
	Resource   *resource.ResourceConfig
	Subsystems []resource.Subsystem
}

func NewCgroupV2Manager(path string) *CgroupV2Manager {
	return &CgroupV2Manager{
		Path:       path,
		Subsystems: fs2.Subsystems,
	}
}

// Apply 将进程pid加入到这个cgroup中
func (c *CgroupV2Manager) Apply(pid int) error {
	for _, subSysIns := range c.Subsystems {
		err := subSysIns.Apply(c.Path, pid)
		if err != nil {
			logrus.Errorf("apply subsystem:%s err:%s", subSysIns.Name(), err)
		}
	}
	return nil
}

// Set 设置cgroup资源限制
func (c *CgroupV2Manager) Set(res *resource.ResourceConfig) error {
	for _, subSysIns := range c.Subsystems {
		err := subSysIns.Set(c.Path, res)
		if err != nil {
			logrus.Errorf("apply subsystem:%s err:%s", subSysIns.Name(), err)
		}
	}
	return nil
}

// Destroy 释放cgroup
func (c *CgroupV2Manager) Destroy() error {
	for _, subSysIns := range c.Subsystems {
		if err := subSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
