package rt

import (
	. "gle/data"
)

type Form struct {
	fn func(*Ctx, []any) (any, error)
}

func (f Form) MetaApply(c *Ctx, args []any) (any, error) {
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

var DECLARE_FORM = Form{func(c *Ctx, args []any) (any, error) {
	arity := len(args)
	if arity != 1 {
		return nil, WrongArity{"declare", arity, []int{1}}
	}
	first := args[0]
	sym, ok := first.(Symbol)
	if !ok {
		return nil, WrongType{"declare", first, "symbol"}
	}
	k := Str(sym)
	v, ok := c.rt.vars[k]
	if !ok {
		v = &variable{}
		c.rt.vars[k] = v
	}
	return nil, nil
}}

var DEFINE_FORM = Form{func(c *Ctx, args []any) (any, error) {
	arity := len(args)
	if arity < 2 {
		return nil, WrongArity{"define", arity, "at least 2"}
	}
	first := args[0]
	sym, ok := first.(Symbol)
	if ok {
		if arity != 2 {
			return nil, WrongArity{"define", arity, []int{2}}
		}

		value, err := c.Eval(args[1])
		if err != nil {
			return nil, err
		}

		c.rt.Define(Str(sym), value)
		return nil, nil
	}

	fnDeclList, ok := first.(List)
	if ok {
		fnNameForm := fnDeclList.First()
		fnName, ok := fnNameForm.(Symbol)
		if !ok {
			return nil, WrongType{"define", fnNameForm, "symbol"}
		}

		argNamesList := fnDeclList.Rest()
		argNames := make([]string, 0, fnDeclList.Count()-1)
		for !argNamesList.IsEmpty() {
			argNameValue := argNamesList.First()
			sym, ok := argNameValue.(Symbol)
			if !ok {
				return nil, WrongType{"define", argNameValue, "symbol"}
			}

			argName := Str(sym)
			argNames = append(argNames, argName)
			argNamesList = argNamesList.Rest()
		}

		body := ListFromSlice(args[1:]).Cons(Symbol{Ns: "", Name: "do"})
		fn := Function{c, argNames, body}
		c.rt.Define(Str(fnName), fn)
		return nil, nil
	}

	return nil, WrongType{"define", first, "symbol or list"}
}}

var IF_FORM = Form{func(c *Ctx, args []any) (any, error) {
	arity := len(args)
	if arity != 2 && arity != 3 {
		return nil, WrongArity{"if", arity, []int{2, 3}}
	}
	cond, err := c.Eval(args[0])
	if err != nil {
		return nil, err
	}

	if Truthy(cond) {
		return c.Eval(args[1])
	}

	if arity == 3 {
		return c.Eval(args[2])
	}

	return nil, nil
}}

var DO_FORM = Form{func(c *Ctx, args []any) (any, error) {
	return doForms(c, ListFromSlice(args))
}}

var LET_FORM = Form{func(c *Ctx, args []any) (any, error) {
	arity := len(args)
	if arity < 1 {
		return nil, WrongArity{"let", 0, "at least 1"}
	}
	bindsForm := args[0]
	binds, ok := bindsForm.(List)
	if !ok {
		return nil, WrongType{"let", bindsForm, "list"}
	}

	if binds.Count()%2 != 0 {
		return nil, WrongArity{"let", "even number of bindings", "uneven number of bindings"}
	}

	bindsMap := make(map[string]any)
	restBinds := binds
	for restBinds.Count() > 0 {
		k := restBinds.First()
		ks, ok := k.(Symbol)
		if !ok {
			return nil, WrongType{"let", k, "symbol"}
		}
		restBinds = restBinds.Rest()
		v := restBinds.First()
		ve, err := c.Extend(bindsMap).Eval(v)
		if err != nil {
			return nil, err
		}
		bindsMap[Str(ks)] = ve
		restBinds = restBinds.Rest()
	}

	if len(bindsMap) > 0 {
		c = c.Extend(bindsMap)
	}

	return doForms(c, ListFromSlice(args[1:]))
}}

var LAMBDA_FORM = Form{func(c *Ctx, args []any) (any, error) {
	arity := len(args)
	if arity < 1 {
		return nil, WrongArity{"lambda", arity, "at least 1"}
	}

	argNamesForm := args[0]
	argNamesList, ok := argNamesForm.(List)
	if !ok {
		return nil, WrongType{"lambda", argNamesForm, "list"}
	}

	argNames := make([]string, 0, argNamesList.Count())

	for !argNamesList.IsEmpty() {
		argNameValue := argNamesList.First()
		sym, ok := argNameValue.(Symbol)
		if !ok {
			return nil, WrongType{"lambda", argNameValue, "symbol"}
		}

		argName := Str(sym)
		argNames = append(argNames, argName)
		argNamesList = argNamesList.Rest()
	}

	body := ListFromSlice(args[1:]).Cons(Symbol{Ns: "", Name: "do"})
	return Function{c, argNames, body}, nil
}}

var STR_FORM = Form{func(c *Ctx, args []any) (any, error) {
    return Str(args...), nil
}}
