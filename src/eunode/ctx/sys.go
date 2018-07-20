package ctx

import "go.uber.org/zap"

var globalContext *System

type System struct {
	Logger *zap.Logger
}

func (sys *System) NewSystem(logger *zap.Logger) {
	sys.Logger = logger
}

func SetupGlobalContext(sysctx *System) {
	globalContext = sysctx
}

func GetGlobalConext() *System {
	return globalContext
}
