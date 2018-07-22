package entry

import (
	"eunode/ctx"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"eunode/core"
	"eunode/launch/launchables"
)

func EunodeEntry() {

	/*
		(1.) Setup Logger
	*/

	logger, error := zap.NewDevelopment()
	logger.Info("(eunode) version 1.0")
	if error != nil {
		fmt.Println("unable to initalize logger context..")
	}
	if unix.Getpid() != 1 {
		logger.Error("(eunode) must start as pid=1, are you starting this from a non init context?")
	}

	/*
		(2.) Setup Context
	*/

	sysCtx := ctx.System{
		Logger: logger,
	}
	ctx.SetupGlobalContext(&sysCtx)

	/*
		(3.) CoreInit, Setup (Basic Init) [FsInit..]
	*/

	core.InvokeLoader()

	/*
		(4.) Launcher (shell, etc..)
	*/

	 shellInst := launchables.Shell{Path:"/bin/mksh"}
	 shellInst.Launch()

}
