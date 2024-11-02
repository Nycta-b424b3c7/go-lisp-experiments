package rt

var PRN_FN = HostFunc{func(args ...any) (any, error) {
	Prn(args...)
	return nil, nil
}}

var PLUS_FN = HostFunc{func(args ...any) (any, error) {
	if len(args) == 0 {
		return 0, nil
	}

	var res any = 0

	for _, arg := range args {
		if i, ok := arg.(int); ok {
			res = res.(int) + i
		} else if f, ok := arg.(float64); ok {
			res = res.(float64) + f
		} else {
			return nil, WrongType{"+", arg, "int/float64"}
		}
	}

	return res, nil
}}

var LT_FN = HostFunc{func(args ...any) (any, error) {
	if len(args) < 2 {
		return 0, WrongArity{"<", len(args), "at least 2"}
	}

	var last float64

	var f float64
	if i, ok := args[0].(int); ok {
		f = float64(i)
	} else if f, ok := args[0].(float64); ok {
		last = f
	} else {
		return nil, WrongType{"<", args[0], "int/float64"}
	}

	last = f

	for _, arg := range args[1:] {
		if i, ok := arg.(int); ok {
			f = float64(i)
		} else if f2, ok := arg.(float64); ok {
			f = f2
		} else {
			return nil, WrongType{"+", arg, "int/float64"}
		}

		if last > f {
			last = f
			continue
		} else {
			return false, nil
		}
	}

	return true, nil
}}

var MINUS_FN = HostFunc{func(args ...any) (any, error) {
	if len(args) == 0 {
		return 0, nil
	}

	var res any = args[0]

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if i, ok := arg.(int); ok {
			res = res.(int) - i
		} else if f, ok := arg.(float64); ok {
			res = res.(float64) - f
		} else {
			return nil, WrongType{"-", arg, "int/float64"}
		}
	}

	return res, nil
}}
