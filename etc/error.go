package etc

import (
	"fmt"
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

