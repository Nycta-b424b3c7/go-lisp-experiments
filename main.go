package main

import (
	"flag"
	"gle/rt"
)

func main() {
    var mainFile string
	flag.StringVar(&mainFile, "main", "", "main file")
	args := flag.CommandLine.Args()
	flag.Parse()

	if mainFile == "" {
		panic("no main file provided")
	}

	err := rt.RunMain(mainFile, args)

	if err != nil {
		panic(err)
	}
}
