package rt

import (
	"fmt"
	. "gle/data"
	. "gle/etc"
)

type NotResolved struct {
	Symbol
}

func (e NotResolved) Error() string {
	return fmt.Sprintf("not resolved: " + e.Symbol.String())
}

type Ctx struct {
	rt    *Rt
	prev  *Ctx
	binds map[string]any
}

func (c *Ctx) Resolve(s Symbol) (any, bool) {
	rt := c.rt
	k := s.Repr()
	for c != nil {
		if v, ok := c.binds[k]; ok {
			return v, true
		}
		c = c.prev
	}
	if v, ok := rt.vars[k]; ok {
		return v.value, true
	}
	return nil, false
}

func (c *Ctx) Extend(binds map[string]any) *Ctx {
	return &Ctx{c.rt, c, binds}
}

func (c *Ctx) EvalList(l List) (any, error) {
	if l.Count() == 0 {
		return EMPTY_LIST, nil
	}
	first := l.First()
	v1, err := c.Eval(first)
	if v1 == nil {
		return nil, InvalidState{Msg: "can't call nil"}
	}
	if err != nil {
		return nil, err
	}
	if ma, ok := v1.(MetaApply); ok {
		return ma.MetaApply(c, l.Rest())
	}
	return nil, InvalidState{Msg: "can't invoke " + Str(v1)}
}

func (c *Ctx) Eval(form any) (any, error) {
	// fmt.Printf("eval %s\n", Str(form))
	if l, ok := form.(List); ok {
		r, err := c.EvalList(l)
		// fmt.Printf(" => %s (%v)\n", Str(r), err)
		return r, err
	}
	if s, ok := form.(Symbol); ok {
		if v, ok := c.Resolve(s); ok {
			// fmt.Printf(" => %s\n", Str(v))
			return v, nil
		}
		return nil, NotResolved{s}
	}
	return form, nil
}
