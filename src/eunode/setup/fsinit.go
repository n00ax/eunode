package setup

import (
	"golang.org/x/sys/unix"
)

type FSInit struct {
	BasePath string
}

func (fsInit *FSInit) Init() {

	/*
		FsInit, initalizes the root filesystem
			(1.) create dev,proc,sys
	*/

	// Create dev, proc, sys
	unix.Mkdir(fsInit.BasePath + "dev", 066)
	unix.Mkdir(fsInit.BasePath + "proc", 066)
	unix.Mkdir(fsInit.BasePath + "sys", 066)

	// Mount procfs and sysfs
	unix.Mount("proc", "/proc", "proc",0, "")
	unix.Mount("sysfs", "/sys", "sysfs", 0, "")

}
