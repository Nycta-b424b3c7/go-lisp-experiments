package rt

import "fmt"

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
	return fmt.Sprintf("wrong argument given to %s: got %s, expected value of type %s", e.Name, e.Got, e.Expected)
}
