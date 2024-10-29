package main

import (
	"fmt"
	"strings"
)

type Repr interface {
	Repr() string
}

func prStr1(val any) string {
	if val == nil {
		return "nil"
	}
	if r, ok := val.(Repr); ok {
		return r.Repr()
	}
	return fmt.Sprintf("%#v", val)
}

func prStr(xs ...any) string {
	strs := make([]string, 0)
	for _, x := range xs {
		strs = append(strs, prStr1(x))
	}
	return strings.Join(strs, " ")

}

func str1(val any) string {
	if val == nil {
		return "nil"
	}
	return fmt.Sprintf("%#v", val)
}

func Str(xs ...any) string {
	var b strings.Builder
	for _, x := range xs {
		b.WriteString(str1(x))
	}
	return b.String()
}

func pr(xs ...any) {
	print(prStr(xs...))
}

func prn(xs ...any) {
	println(prStr(xs...))
}
