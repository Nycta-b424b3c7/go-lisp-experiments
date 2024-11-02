package rt

import (
	"gle/reader"
	"os"
	"path"
	"strings"
)

var coreBinds = map[string]any{
	"declare": DECLARE_FORM,
	"define":  DEFINE_FORM,
	"if":      IF_FORM,
	"do":      DO_FORM,
	"let":     LET_FORM,
	"lambda":  LAMBDA_FORM,

	"+":   PLUS_FN,
	"prn": PRN_FN,
}

type variable struct {
	bound bool
	value any
}

func (v *variable) Deref() (any, error) {
	if v.bound {
		return v.value, nil
	} else {
		return nil, NOT_BOUND
	}
}

func getPath() []string {
	var path []string
	var pathEnv = os.Getenv("GLE_PATH")
	if pathEnv == "" {
		pathEnv = "."
	}
	path = strings.Split(pathEnv, ";")
	return path
}

type Rt struct {
	path []string
	vars map[string]*variable
}

func NewRt() *Rt {
	return &Rt{getPath(), make(map[string]*variable)}
}

func (rt *Rt) Declare(symRepr string) {
	v, ok := rt.vars[symRepr]
	if !ok {
		v = &variable{}
		rt.vars[symRepr] = v
	}
}

func (rt *Rt) Define(symRepr string, value any) {
	rt.Declare(symRepr)
	v := rt.vars[symRepr]
	v.bound = true
	v.value = value
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

func (rt *Rt) require(ns string) error {
	file := strings.Replace(ns, ".", "/", -1) + ".gle"
	return rt.load(file)
}

func (rt *Rt) load(file string) error {
	var forms []any
	var err error

	for _, p := range rt.path {
		f, err := os.Open(path.Join(p, file))
		if err != nil {
			break
		}

		i, err := f.Stat()
		if err != nil {
			break
		}

		if i.IsDir() {
			break
		}

		forms, err = reader.Read(f)
		break
	}

	if err != nil {
		return err
	}

	if forms == nil {
		return LoadError{FileName: file, Path: rt.path}
	}

	for _, f := range forms {
		_, err := rt.Eval(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func Truthy(form any) bool {
	return form != nil && form != false
}

func RunMain(mainNs string, arguments []string) error {
	rt := NewRt()

	err := rt.require(mainNs)
	if err != nil {
		return err
	}

	mainFn, err := rt.evalStr(mainNs + "/main")
	if err != nil {
		return err
	}

	if a, ok := mainFn.(MetaApply); ok {
		anyArgs := make([]any, 0, len(arguments))
		for _, anyArg := range arguments {
			anyArgs = append(anyArgs, anyArg)
		}
		_, err := a.MetaApply(rt.Ctx(), anyArgs)
		return err
	} else {
		return WrongType{"gle", mainFn, "function"}
	}
}
