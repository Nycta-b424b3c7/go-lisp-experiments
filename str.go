package main

import (
	"fmt"
	"strings"
)

func str1(val any) string {
	if val == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", val)
}

func str(xs ...any) string {
	var b strings.Builder
	for _, x := range xs {
		b.WriteString(str1(x))
	}
	return b.String()
}

func pr(xs ...any) {
	print(str(xs...))
}

func prn(xs ...any) {
	println(str(xs...))
}
