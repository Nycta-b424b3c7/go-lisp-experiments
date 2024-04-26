package rt

import (
	. "gle/data"
)

type MetaApply interface {
	MetaApply(*Ctx, List) (any, error)
}

type MetaFunction struct {
	Ctx      *Ctx
	ArgNames List
	Body     List
}

func (mf MetaFunction) MetaApply(c *Ctx, args List) (any, error) {
	args_binds := make(map[string]any)
	al := args
	an := mf.ArgNames
	for {
		k := an.First()
		v := al.First()
		if k == nil {
			break
		}
		ks, ok := k.(Symbol)
		if !ok {
			panic("meta function definition parameters are supposed to be symbols")
		}
		args_binds[ks.Repr()] = v
		al = al.Rest()
		an = an.Rest()
	}
	c1 := mf.Ctx.Extend(args_binds)
	return c1.Eval(mf.Body)
}
