package main

var PRN_FN = HostFunc{func(args ...any) (any, error) {
	prn(args...)
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
