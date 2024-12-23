package reader

import (
	. "gle/data"
	"io"
	"strconv"
	"strings"
)

type gleParser struct {
	l *lexer
}

func newGleParser(r io.Reader) *gleParser {
	l := newLexer(r)
	go l.read()
	return &gleParser{l}
}

func inferAtom(tok string) (any, error) {
	btok := []byte(tok)
	if len(tok) == 0 {
		panic("inferAtom / len(tok) == 0")
	}
	if tok == "nil" {
		return nil, nil
	}
	if tok == "true" {
		return true, nil
	}
	if tok == "false" {
		return false, nil
	}
	if intRe.Match(btok) {
		return strconv.Atoi(tok)
	}
	if hexRe.Match(btok) {
		return strconv.ParseInt(tok, 16, 64)
	}
	if floatRe.Match(btok) {
		return strconv.ParseFloat(tok, 64)
	}
	if tok[0] == ':' && len(tok) > 1 {
		return parseKeyword(tok)
	}
	return parseSymbol(tok)
}

func (r *gleParser) nextList() (List, bool, error) {
	forms := make([]any, 0)
	for {
		form, ok, err := r.nextForm1(true)
		if err != nil {
			return EMPTY_LIST, false, err
		}
		if !ok {
			break
		}
		forms = append(forms, form)
	}
	list := ListFromSlice(forms)
	return list, true, nil
}

func (r *gleParser) nextForm1(inList bool) (any, bool, error) {
	tok, ok := <-r.l.tokens
	if !ok && inList {
		return nil, false, UNEXPECTED_EOF
	}
	if tok == ")" && inList {
		return nil, false, nil
	}
	if tok == ")" && !inList {
		return nil, false, UnexpectedToken{tok}
	}
	if tok == "(" {
		return r.nextList()
	}
	if ok {
		atom, err := inferAtom(tok)
		return atom, err == nil, err
	}
	return nil, false, nil
}

func (r *gleParser) nextForm() (any, bool, error) {
	return r.nextForm1(false)
}

type ReadResult struct {
	Form any
	Err  error
}

func Read(r io.Reader) ([]any, error) {
	forms := make([]any, 0)
	r1 := newGleParser(r)
	for {
		form, ok, err := r1.nextForm()
		if err != nil {
			return nil, err
		} else if ok {
			forms = append(forms, form)
		} else {
			break
		}
	}
	return forms, nil
}

func ReadString(s string) ([]any, error) {
	r := strings.NewReader(s)
	return Read(r)
}
