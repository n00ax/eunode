package launchables

import (
	"golang.org/x/sys/unix"
	"os"
)

type Shell struct{
	Path string
}

func (shell *Shell) Launch(){

	unix.Exec("/bin/mksh", []string{}, os.Environ())

}