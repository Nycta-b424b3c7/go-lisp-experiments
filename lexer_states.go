package main

import (
	"bufio"
	"io"
	"unicode"
)

func next(r *bufio.Reader) (rune, bool) {
	c, _, err := r.ReadRune()
	if err == io.EOF {
		return 0, false
	}
	if err != nil {
		panic(err) // TODO actual error handling
	}
	return c, true
}

func back(r *bufio.Reader) {
	if err := r.UnreadRune(); err != nil {
		panic(err) // TODO actual error handling
	}
}

func isSpc(c rune) bool {
	return unicode.IsSpace(c)
}

func isDelim(c rune) bool {
	return c == '(' || c == ')'
}

func readAtom(l *lexer) lexerState {
	r := l.reader
	ts := l.tokens
	cs := make([]rune, 0)
	for {
		c, ok := next(r)
		if !ok {
			break
		}
		if isSpc(c) {
			break
		}
		if isDelim(c) {
			back(r)
			break
		}
		cs = append(cs, c)
	}
	if len(cs) > 0 {
		ts <- string(cs)
	}
	return lexerSwitch
}

func lexerSwitch(l *lexer) lexerState {
	r := l.reader
	ts := l.tokens
	c, ok := next(r)
	if !ok {
		return nil
	} else if c == '(' {
		ts <- "("
		return lexerSwitch
	} else if c == ')' {
		ts <- ")"
		return lexerSwitch
	} else if unicode.IsSpace(c) {
		return skipWhitespace
	} else {
		back(r)
		return readAtom
	}
}

func skipWhitespace(l *lexer) lexerState {
	r := l.reader
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			panic(err) // TODO actual error handling
		}
		if !unicode.IsSpace(c) {
			back(r)
			break
		}
	}
	return lexerSwitch
}
