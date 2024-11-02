package main

import (
	"flag"
	"gle/rt"
)

func main() {

	mainFile := flag.String("main", "", "main file")
	args := flag.CommandLine.Args()
	flag.Parse()

	if mainFile == nil {
		panic("no main file provided")
	}

	err := rt.RunMain(*mainFile, args)

	if err != nil {
		panic(err)
	}
}
