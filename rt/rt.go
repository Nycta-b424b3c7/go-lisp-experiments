package rt

import (
	"gle/reader"
)

var coreBinds = map[string]any{
	"def":     DEF_FORM,
	"if":      IF_FORM,
	"do":      DO_FORM,
	"let":     LET_FORM,
	"meta-fn": META_FN_FORM,
	"fn":      FN_FORM,
}

type variable struct {
	value any
}

type Rt struct {
	vars map[string]*variable
}

func NewRt() *Rt {
	return &Rt{make(map[string]*variable)}
}

func (rt *Rt) Ctx() *Ctx {
	return &Ctx{rt, nil, coreBinds}
}

func (rt *Rt) Eval(form any) (any, error) {
	return rt.Ctx().Eval(form)
}

func (rt *Rt) evalStr(s string) (any, error) {
	f, err := reader.ReadString(s)
	if err != nil {
		return nil, err
	}
	if len(f) != 1 {
		return nil, WrongArity{"evalStr", len(f), []int{1}}
	}
	return rt.Eval(f[0])
}

func Truthy(form any) bool {
	return form != nil && form != false
}
