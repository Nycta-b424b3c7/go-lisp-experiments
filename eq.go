package main

import "reflect"

type Eq interface {
	Eq(other any) bool
}

func IsEq(a, b any) bool {
	if ae, ok := a.(Eq); ok {
		return ae.Eq(b)
	}
	return a == b || reflect.DeepEqual(a, b)
}
