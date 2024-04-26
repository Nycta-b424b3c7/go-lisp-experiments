package rt

import "testing"

func TestDef(t *testing.T) {
	rt := NewRt()
	var r any
	var err error
	r, err = rt.evalStr(`(def x 1)`)
	if err != nil {
		t.Fail()
	}
	if r != nil {
		t.Fail()
	}

	r, err = rt.evalStr(`x`)
	if err != nil {
		t.Fail()
	}
	if r != 1 {
		t.Fail()
	}

	r, err = rt.evalStr(`(def x 2)`)
	if err != nil {
		t.Fail()
	}
	if r != nil {
		t.Fail()
	}

	r, err = rt.evalStr(`x`)
	if err != nil {
		t.Fail()
	}
	if r != 2 {
		t.Fail()
	}
}

func TestIf(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(if true 1 0)`)
	if err != nil {
		t.Fail()
	}
	if r != 1 {
		t.Fail()
	}
}

func TestDo(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(do 1 2 3)`)
	if err != nil {
		t.Fail()
	}
	if r != 3 {
		t.Fail()
	}
}

func TestLet(t *testing.T) {
	rt := NewRt()
	r, err := rt.evalStr(`(let (x 1 y 2 z 3) x y z)`)
	if err != nil {
		t.Fail()
	}
	if r != 3 {
		t.Fail()
	}
}
