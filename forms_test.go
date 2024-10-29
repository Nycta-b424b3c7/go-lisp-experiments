package main

import "testing"

func TestDef(t *testing.T) {
	rt := NewRt()
	var r any
	var err error
	r, err = rt.evalStr(`(def x 1)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != nil {
		t.FailNow()
	}

	r, err = rt.evalStr(`x`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 1 {
		t.FailNow()
	}

	r, err = rt.evalStr(`(def x 2)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != nil {
		t.FailNow()
	}

	r, err = rt.evalStr(`x`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 2 {
		t.FailNow()
	}
}

func TestIf(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(if true 1 0)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 1 {
		t.FailNow()
	}
}

func TestDo(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(do 1 2 3)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.FailNow()
	}
}

func TestLet(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(let (x 1 y 2 z 3) x y z)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.FailNow()
	}
}

func TestLambda(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`((lambda (x y) (+ x y)) 1 2)`)
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.FailNow()
	}
}
