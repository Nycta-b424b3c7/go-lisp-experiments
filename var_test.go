package main

import "testing"

func TestVars(t *testing.T) {
	rt := NewRt()
	var r any
	var err error

	r, err = rt.evalStr(`(declare x)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != nil {
		t.Fail()
	}

	r, err = rt.evalStr(`(define (more) (+ x 1))`)
	if err != nil {
		t.Fatal(err)
	}
	if r != nil {
		t.Fail()
	}

	r, err = rt.evalStr(`(define x 1)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != nil {
		t.Fail()
	}

	r, err = rt.evalStr(`(more)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 2 {
		t.Fail()
	}
}
