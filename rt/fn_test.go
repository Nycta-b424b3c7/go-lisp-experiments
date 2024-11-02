package rt

import (
	"testing"
)

func TestFn(t *testing.T) {
	rt := NewRt()

	rt.vars["="] = &variable{true, HostFunc{func(args ...any) (any, error) {
		if len(args) < 1 {
			return nil, WrongArity{"=", len(args), "at least 1"}
		}
		x := args[0]
		for _, y := range args[1:] {
			if !IsEq(x, y) {
				return false, nil
			}
			x = y
		}
		return true, nil
	}}}

	rt.vars["mod"] = &variable{true, HostFunc{func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, WrongArity{"mod", len(args), 2}
		}

		a := args[0]
		b := args[1]

		x, ok := a.(int)
		if !ok {
			return nil, WrongType{"mod", a, "int"}
		}

		y, ok := b.(int)
		if !ok {
			return nil, WrongType{"mod", b, "int"}
		}

		r := x % y
		return r, nil
	}}}

	_, err := rt.evalStr(`
        (def even? (fn (x) (= 0 (mod x 2)))))
        `)
	if err != nil {
		t.Log("fail 1")
		t.Fatal(err)
	}

	r1, err := rt.evalStr(`(even? 1)`)
	if err != nil {
		t.Log("fail 2")
		t.Fatal(err)
	}

	if r1 != false {
		t.Logf("%+v\n", r1)
		t.Fail()
	}

	r2, err := rt.evalStr(`(even? 2)`)
	if err != nil {
		t.Logf("%+v\n", r1)
		t.Fatal(err)
	}

	if r2 != true {
		t.Fail()
	}
}
