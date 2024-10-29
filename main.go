package main

import (
	"flag"
)

func main() {

	mainFile := flag.String("main", "", "main file")
	args := flag.CommandLine.Args()
	flag.Parse()

	if mainFile == nil {
		panic("no main file provided")
	}

	err := RunMain(*mainFile, args)

	if err != nil {
		panic(err)
	}
}
