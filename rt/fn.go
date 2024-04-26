package rt

import (
	. "gle/data"
)

type HostFunc struct {
	fn func(args ...any) (any, error)
}

func (hf HostFunc) MetaApply(c *Ctx, al List) (any, error) {
	args := make([]any, 0)
	for al.Count() > 0 {
		v, err := c.Eval(al.First())
		if err != nil {
			return nil, err
		}
		args = append(args, v)
		al = al.Rest()
	}
	return hf.fn(args...)
}

type Function struct {
	Ctx      *Ctx
	ArgNames List
	Body     List
}

func (f Function) MetaApply(c *Ctx, args List) (any, error) {
	args_binds := make(map[string]any)
	al := args
	an := f.ArgNames
	for {
		k := an.First()
		if k == nil {
			break
		}
		ks, ok := k.(Symbol)
		if !ok {
			panic("meta function definition parameters are supposed to be symbols")
		}
		v, err := c.Eval(al.First())
		if err != nil {
			return nil, err
		}
		args_binds[ks.Repr()] = v
		al = al.Rest()
		an = an.Rest()
	}
	c1 := f.Ctx.Extend(args_binds)
	return c1.Eval(f.Body)
}
