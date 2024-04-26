package etc

import "fmt"

type InvalidState struct {
	Msg string
}

func (e InvalidState) Error() string {
	return fmt.Sprintf("invalid state: %s", e.Msg)
}
