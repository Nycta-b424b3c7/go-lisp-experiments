package reader

import "fmt"

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

