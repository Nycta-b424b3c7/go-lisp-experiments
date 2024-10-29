package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	src := ``
	p := newGleParser(strings.NewReader(src))
	forms := make([]any, 0)
	exp := []any{}
	for {
		form, ok, err := p.nextForm()
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			break
		}
		forms = append(forms, form)
	}
	for i := 0; i < min(len(forms), len(exp)); i++ {
		a := forms[i]
		b := exp[i]
		if a != b {
			t.Logf("%+v\n", a)
			t.Logf("%+v\n", b)
			t.Fail()
		}
	}
}
