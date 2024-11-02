package reader

import (
	. "gle/data"
	"regexp"
	"strings"
)

func re(s string) *regexp.Regexp {
	p, err := regexp.Compile(s)
	if err != nil {
		panic(err)
	}
	return p
}

var intRe = re(`^[+-]?[0-9]+$`)
var hexRe = re(`^[+-]?0x[0-9a-fA-F]+$`)
var floatRe = re(`^[+-]?\d+\.\d+$`)

func parseSymbol(tok string) (Symbol, error) {
	var err error
	var ns, name string
	if tok == "/" {
		name = "/"
	} else if strings.ContainsRune(tok, '/') {
		parts := strings.SplitN(tok, "/", 2)
		ns, name = parts[0], parts[1]
		if len(ns) == 0 {
			err = InvalidSymbol{Sym: tok}
		}
		if name != "/" && strings.ContainsRune(name, '/') {
			err = InvalidSymbol{Sym: tok}
		}
	} else {
		name = tok
	}
	return Symbol{Ns: ns, Name: name}, err
}

func parseKeyword(tok string) (Keyword, error) {
	var err error
	var sym Symbol
	if len(tok) < 2 {
		err = InvalidKeyword{Sym: tok}
	}
	sym, err = parseSymbol(tok[1:])
	if err != nil {
		err = InvalidKeyword{Sym: tok}
	}

	return Keyword{Sym: &sym}, err
}
