package rt

import (
	"fmt"
	. "gle/data"
)

type NotResolved struct {
	Symbol
}

func (e NotResolved) Error() string {
	return fmt.Sprintf("not resolved: " + Str(e.Symbol))
}

type Ctx struct {
	rt    *Rt
	outer *Ctx
	binds map[string]any
}

func (c *Ctx) Resolve(s Symbol) (*variable, bool) {
	rt := c.rt
	k := Str(s)
	for c != nil {
		if v, ok := c.binds[k]; ok {
			return &variable{true, v}, true
		}
		c = c.outer
	}
	if v, ok := rt.vars[k]; ok {
		return v, true
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

	if err != nil {
		return nil, err
	}
	if v1 == nil {
		return nil, InvalidState{Msg: "can't invoke nil"}
	}
	if ma, ok := v1.(MetaApply); ok {
		return ma.MetaApply(c, l.Rest().ToSlice())
	}
	return nil, InvalidState{Msg: "can't invoke " + Str(v1)}
}

func (c *Ctx) Eval(form any) (any, error) {
	if s, ok := form.(Symbol); ok {
		if v, ok := c.Resolve(s); ok {
			return v.Deref()
		}

		return nil, NotResolved{s}
	}
	if l, ok := form.(List); ok {
		r, err := c.EvalList(l)
		return r, err
	}
	return form, nil
}
