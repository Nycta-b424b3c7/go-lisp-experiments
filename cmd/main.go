package main

import (
	"gle/rt"
)

func main() {
	err := rt.Run()
	if err != nil {
		panic(err)
	}
}
