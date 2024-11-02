package rt

import (
    . "gle/data"
)

type MetaApply interface {
	MetaApply(*Ctx, []any) (any, error)
}

type MetaFunction struct {
	Ctx      *Ctx
	ArgNames []string
	Body     List
}

func (mf MetaFunction) MetaApply(c *Ctx, args []any) (any, error) {
	argBinds := make(map[string]any)

	for i := 0; i < len(mf.ArgNames); i++ {
		k := mf.ArgNames[i]
		var v any = nil
		if i < len(args) {
			v = args[i]
		}
		argBinds[k] = v
	}

	c1 := mf.Ctx.Extend(argBinds)
	return c1.Eval(mf.Body)
}
