package rt

import (
	. "gle/data"
)

type Form struct {
	fn func(*Ctx, List) (any, error)
}

func (f Form) MetaApply(c *Ctx, args List) (any, error) {
	return f.fn(c, args)
}

func doForms(c *Ctx, forms List) (any, error) {
	var res any
	var err error
	for forms.Count() > 0 {
		res, err = c.Eval(forms.First())
		if err != nil {
			return nil, err
		}
		forms = forms.Rest()
	}
	return res, err
}

var DEF_FORM = Form{func(c *Ctx, args List) (any, error) {
	arity := args.Count()
	if arity != 2 {
		return nil, WrongArity{"def", arity, []int{2}}
	}
	first := args.First()
	sym, ok := first.(Symbol)
	if !ok {
		return nil, WrongType{"def", first, "symbol"}
	}
	k := sym.Repr()
	v, ok := c.rt.vars[k]
	if !ok {
		v = &variable{}
		c.rt.vars[k] = v
	}
	value, err := c.Eval(args.Rest().First())
	if err != nil {
		return nil, err
	}
	v.value = value
	return nil, nil
}}

var IF_FORM = Form{func(c *Ctx, args List) (any, error) {
	arity := args.Count()
	if arity != 2 && arity != 3 {
		return nil, WrongArity{"if", arity, []int{2, 3}}
	}
	cond, err := c.Eval(args.First())
	if err != nil {
		return nil, err
	}
	rest := args.Rest()
	second := rest.First()
	if Truthy(cond) {
		return c.Eval(second)
	}
	last := rest.Rest().First()
	if last != nil {
		return c.Eval(last)
	}

	return nil, nil
}}

var DO_FORM = Form{func(c *Ctx, args List) (any, error) {
	return doForms(c, args)
}}

var LET_FORM = Form{func(c *Ctx, args List) (any, error) {
	arity := args.Count()
	if arity < 1 {
		return nil, WrongArity{"let", 0, "at least 1"}
	}
	binds_a := args.First()
	binds, ok := binds_a.(List)
	if !ok {
		return nil, WrongType{"let", binds_a, "list"}
	}

	if binds.Count()%2 != 0 {
		return nil, WrongArity{"let", "even number of bindings", "uneven number of bindings"}
	}

	binds_m := make(map[string]any)
	rest_binds := binds
	for rest_binds.Count() > 0 {
		k := rest_binds.First()
		ks, ok := k.(Symbol)
		if !ok {
			return nil, WrongType{"let", k, "symbol"}
		}
		rest_binds = rest_binds.Rest()
		v := rest_binds.First()
		ve, err := c.Extend(binds_m).Eval(v)
		if err != nil {
			return nil, err
		}
		binds_m[ks.Repr()] = ve
		rest_binds = rest_binds.Rest()
	}

	if len(binds_m) > 0 {
		c = c.Extend(binds_m)
	}

	return doForms(c, args.Rest())
}}

var META_FN_FORM = Form{func(c *Ctx, args List) (any, error) {
	arity := args.Count()
	if arity < 1 {
		return nil, WrongArity{"meta-fn", arity, "at least 1"}
	}

	arg_names_a := args.First()
	arg_names, ok := arg_names_a.(List)
	if !ok {
		return nil, WrongType{"meta-fn", arg_names_a, "list"}
	}

	validation_err := Reduce(func(err error, item any) error {
		if err != nil {
			return err
		}

		_, ok := item.(Symbol)
		if !ok {
			return WrongType{"meta-fn", item, "list"}
		}

		return nil
	}, nil, arg_names)

	if validation_err != nil {
		return nil, validation_err
	}

	body := args.Rest()
	return MetaFunction{c, arg_names, body}, nil
}}

var FN_FORM = Form{func(c *Ctx, args List) (any, error) {
	arity := args.Count()
	if arity < 1 {
		return nil, WrongArity{"meta-fn", arity, "at least 1"}
	}

	arg_names_a := args.First()
	arg_names, ok := arg_names_a.(List)
	if !ok {
		return nil, WrongType{"meta-fn", arg_names_a, "list"}
	}

	validation_err := Reduce(func(err error, item any) error {
		if err != nil {
			return err
		}

		_, ok := item.(Symbol)
		if !ok {
			return WrongType{"meta-fn", item, "list"}
		}

		return nil
	}, nil, arg_names)

	if validation_err != nil {
		return nil, validation_err
	}

	body := args.Rest().Cons(Symbol{Ns: "", Name: "do"})
	return Function{c, arg_names, body}, nil
}}
