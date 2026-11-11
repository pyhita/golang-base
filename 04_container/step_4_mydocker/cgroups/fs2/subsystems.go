package fs2

import (
	"github.com/pyhita/mydocker/cgroups/resource"
)

var Subsystems = []resource.Subsystem{
	&CpusetSubSystem{},
	&MemorySubSystem{},
	&CpuSubSystem{},
}
