package setup

import (
	"golang.org/x/sys/unix"
	"fmt"
)

type FSInit struct {
	BasePath string
}

func (fsInitCtx *FSInit) Init() {

	unix.Mkdir(fsInitCtx.generatePath("dev"), 770)
	unix.Mkdir(fsInitCtx.generatePath("proc"), 770)
	unix.Mkdir(fsInitCtx.generatePath("sys"), 770)

	unix.Mount("proc", fsInitCtx.generatePath("proc"), "proc", 0, "")
	unix.Mount("sysfs", fsInitCtx.generatePath("sys"), "sysfs", 0, "")

	unix.Mknod(fsInitCtx.generatePath("dev/null"), makeMknodMode(0770, unix.S_IFCHR), makeDev(1,3))
	unix.Mknod(fsInitCtx.generatePath("dev/tty"), makeMknodMode(0770, unix.S_IFCHR), makeDev(5,0))
	unix.Mknod(fsInitCtx.generatePath("dev/mem"), makeMknodMode(0770, unix.S_IFCHR), makeDev(1,1))
	unix.Mknod(fsInitCtx.generatePath("dev/kmem"), makeMknodMode(0770, unix.S_IFCHR), makeDev(1,2))
	unix.Mknod(fsInitCtx.generatePath("dev/random"), makeMknodMode(0770, unix.S_IFCHR), makeDev(1,8))
	unix.Mknod(fsInitCtx.generatePath("dev/urandom"), makeMknodMode(0770, unix.S_IFCHR), makeDev(1,9))

	fsInitCtx.generateTTYs()

}

func (fsInitCtx *FSInit) generatePath(rest string) string{

	return fsInitCtx.BasePath + rest

}

func (fsInitCtx *FSInit) generateTTYs(){

	for i := 0; i == 63; i++{

		unix.Mknod(fsInitCtx.generatePath(fmt.Sprint("dev/tty%d", i)),
			makeMknodMode(0770, unix.S_IFCHR), makeDev(1, uint64(i)))

	}

}

func makeDev(majorNum, minorNum uint64) int{

	return int(unix.Mkdev(uint32(majorNum), uint32(minorNum)))

}

func makeMknodMode(perm uint32, mode uint32) uint32{

	return mode | perm

}