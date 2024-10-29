package main

import (
	"fmt"
	"path"
	"strings"
)

type NotImplemented struct{}

var NOT_IMPLEMENTED = NotImplemented{}

func (_ NotImplemented) Error() string {
	return "not implemented"
}

type InvalidState struct {
	Msg string
}

func (e InvalidState) Error() string {
	return fmt.Sprintf("invalid state: %s", e.Msg)
}

type WrongArity struct {
	Name     string
	Got      any
	Expected any
}

func (e WrongArity) Error() string {
	return fmt.Sprintf("wrong arity given to %v: got %v, expected %v", e.Name, e.Got, e.Expected)
}

type WrongType struct {
	Name     string
	Got      any
	Expected string
}

func (e WrongType) Error() string {
	return fmt.Sprintf("wrong argument given to %s: got %#v (%T), expected value of type %s", e.Name, e.Got, e.Got, e.Expected)
}

type UnexpectedEof struct{}

var UNEXPECTED_EOF = UnexpectedEof{}

func (e UnexpectedEof) Error() string {
	return "UnexpectedEof"
}

type UnexpectedToken struct {
	tok string
}

func (e UnexpectedToken) Error() string {
	return fmt.Sprintf("UnexpectedToken: %s", e.tok)
}

type NotBound struct{}

var NOT_BOUND = NotBound{}

func (_ NotBound) Error() string {
	return "variable not bound"
}

type LoadError struct {
	FileName string
	Path     []string
}

func (e LoadError) Error() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("file %s not found on $GLE_PATH:\n", e.FileName))

	for _, pathDir := range e.Path {
		b.WriteString(fmt.Sprintf("\t%s\n", path.Join(pathDir, e.FileName)))
	}

	return b.String()
}
