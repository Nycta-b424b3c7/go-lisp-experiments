package rt

import (
	. "gle/data"
)

type HostFunc struct {
	fn func(...any) (any, error)
}

func (hf HostFunc) MetaApply(c *Ctx, argForms []any) (any, error) {
	args := make([]any, len(argForms))

	for i := 0; i < len(argForms); i++ {
		v, err := c.Eval(argForms[i])
		if err != nil {
			return nil, err
		}
		args[i] = v
	}

	return hf.fn(args...)
}

type Function struct {
	Ctx      *Ctx
	ArgNames []string
	Body     List
}

func (f Function) MetaApply(c *Ctx, argForms []any) (any, error) {
	argBinds := make(map[string]any)

	args := make([]any, len(argForms))

	for i := 0; i < len(argForms); i++ {
		v, err := c.Eval(argForms[i])
		if err != nil {
			return nil, err
		}
		args[i] = v
	}

	for i := 0; i < len(f.ArgNames); i++ {
		k := f.ArgNames[i]
		var v any = nil
		if i < len(args) {
			v = args[i]
		}
		argBinds[k] = v
	}

	c1 := f.Ctx.Extend(argBinds)
	return c1.Eval(f.Body)
}
