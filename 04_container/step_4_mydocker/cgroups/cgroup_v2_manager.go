package cgroups

import (
	"github.com/pyhita/mydocker/cgroups/resource"
)

type CgroupV2Manager struct {
}

func (c CgroupV2Manager) Apply(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (c CgroupV2Manager) Set(res *resource.ResourceConfig) error {
	//TODO implement me
	panic("implement me")
}

func (c CgroupV2Manager) Destroy() error {
	//TODO implement me
	panic("implement me")
}

func NewCgroupV2Manager(path string) *CgroupV2Manager {
	return &CgroupV2Manager{}
}
