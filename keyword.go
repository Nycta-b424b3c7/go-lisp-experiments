package main

import (
	"fmt"
)

type InvalidKeyword struct {
	Sym string
}

func (e InvalidKeyword) Error() string {
	return fmt.Sprintf("Invalid keyword: %s", e.Sym)
}

type Keyword struct {
	Sym *Symbol
}

func (k Keyword) Repr() string {
    return ":" + k.Sym.Repr()
}

func (k Keyword) String() string {
	return k.Repr()
}
