package reader

import (
	"bufio"
	"io"
)

type lexer struct {
	reader *bufio.Reader
	tokens chan string
}

func newLexer(r io.Reader) *lexer {
	r1 := bufio.NewReader(r)
	toks := make(chan string)
	return &lexer{r1, toks}
}

type lexerState func(*lexer) lexerState

func (l *lexer) read() {
	var state lexerState = lexerSwitch
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}
