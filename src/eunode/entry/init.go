package entry

import (
	"eunode/ctx"
	"eunode/setup"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"eunode/launch"
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

	fsInitCtx := setup.FSInit{
		BasePath: "/",
	}
	fsInitCtx.Init()

	/*
		(4.) Basic Launcher
	 */
	 shellCtx := launch.Shell{

	 }
	 shellCtx.Launch()

}
