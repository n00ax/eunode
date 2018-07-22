package main

import (
	"eunode/entry"
	"time"
)

func main() {

	/*
		Bootstrap Entry Point
	*/

	entry.EunodeEntry()

	time.Sleep(time.Second*10000)
	
}
