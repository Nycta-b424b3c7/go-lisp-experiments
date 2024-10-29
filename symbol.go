package main

import (
	"fmt"
)

type InvalidSymbol struct {
	Sym string
}

func (e InvalidSymbol) Error() string {
	return fmt.Sprintf("Invalid symbol: %s", e.Sym)
}

type Symbol struct {
	Ns, Name string
}

func (s Symbol) Repr() string {
	if len(s.Ns) > 0 {
		return s.Ns + "/" + s.Name
	} else {
		return s.Name
	}
}

func (s Symbol) String() string {
	return s.Repr()
}
