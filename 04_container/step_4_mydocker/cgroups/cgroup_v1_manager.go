package cgroups

import "github.com/pyhita/mydocker/cgroups/resource"

type CgroupV1Manager struct {
}

func (c CgroupV1Manager) Apply(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (c CgroupV1Manager) Set(res *resource.ResourceConfig) error {
	//TODO implement me
	panic("implement me")
}

func (c CgroupV1Manager) Destroy() error {
	//TODO implement me
	panic("implement me")
}

func NewCgroupV1Manager(path string) *CgroupV1Manager {
	return &CgroupV1Manager{}
}
