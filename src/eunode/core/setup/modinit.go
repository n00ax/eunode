package setup

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

type ModInit struct{

	ModDir string

}

func (modCtx *ModInit) Init(){



}

func (modCtx *ModInit) load(){



}

func (modCtx *ModInit) hookFinitModule(fd uint32, paramValues *string, flags int){

	unix.Syscall(unix.SYS_FINIT_MODULE, uintptr(fd), uintptr(unsafe.Pointer(paramValues)), uintptr(flags))

}