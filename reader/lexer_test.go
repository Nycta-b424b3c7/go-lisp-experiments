package reader

import (
	"slices"
	"strings"
	"testing"
)

func TestLexerEmpty(t *testing.T) {
	src := ""
	exp := []string{}
	l := newLexer(strings.NewReader(src))
	ts := make([]string, 0)
	go l.read()
	for t := range l.tokens {
		ts = append(ts, t)
	}

	if !slices.Equal(ts, exp) {
		t.Fail()
	}
}

func TestLexer1(t *testing.T) {
	src := "(asd (qwe zxc))"
	exp := []string{"(", "asd", "(", "qwe", "zxc", ")", ")"}
	l := newLexer(strings.NewReader(src))
	ts := make([]string, 0)
	go l.read()
	for t := range l.tokens {
		ts = append(ts, t)
	}
	t.Logf("ts:  %+v\n", ts)
	t.Logf("exp: %+v\n", exp)
	if len(ts) != len(exp) {
		t.Fail()
	}
	for i := 0; i < len(exp); i++ {
		a := ts[i]
		b := exp[i]
		if a != b {
			t.Logf("%+v\n", a)
			t.Logf("%+v\n", b)
			t.Fail()
		}
	}
}

func TestLexer2(t *testing.T) {
	src := `
        (defn fibonacci (n)
          (if (< n 1)
            1
            (fibonnaci (- n 1) (- n 2))))
    `
	exp := []string{"(", "defn", "fibonacci", "(", "n", ")", "(", "if", "(", "<", "n", "1", ")", "1", "(", "fibonnaci", "(", "-", "n", "1", ")", "(", "-", "n", "2", ")", ")", ")", ")"}
	l := newLexer(strings.NewReader(src))
	ts := make([]string, 0)
	go l.read()
	for t := range l.tokens {
		ts = append(ts, t)
	}
	t.Logf("ts:  %+v\n", ts)
	t.Logf("exp: %+v\n", exp)
	if len(ts) != len(exp) {
		t.Fail()
	}
	for i := 0; i < len(exp); i++ {
		a := ts[i]
		b := exp[i]
		if a != b {
			t.Logf("%+v %+v\n", a, b)
			t.Fail()
		}
	}
}
