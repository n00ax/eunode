package core

import "eunode/core/setup"

func InvokeLoader(){

	loaders := [...]setup.Init{

		&setup.FSInit{BasePath: "/"},
		&setup.ModInit{ModDir : "/system/modules/"},

	}

	for e := range loaders {

		loaders[e].Init()

	}

}