package rt

import (
	"fmt"
	"gle/data"
	"gle/reader"
	"os"
	"path"
	"strings"
)

var Path []string = strings.Split(os.Getenv("RUV_PATH"), ";")

type LoadError struct {
	Mod string
}

func (e LoadError) Error() string {
	return fmt.Sprintf("module %s not found", e.Mod)
}

func load(mod string) error {
	file := mod + ".gle"
	var forms []any
	var err error
	for _, p := range Path {
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
		return LoadError{mod}
	}

	for _, f := range forms {
		_, err := NewRt().Eval(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func Run() error {
	if len(os.Args) < 2 {
		return WrongArity{"gle", len(os.Args) - 1, "at least 1"}
	}

	main_mod := os.Args[1]
	args := os.Args[2:]

	err := load(main_mod)
	if err != nil {
		return err
	}

	rt := NewRt()

	main_fn, err := rt.evalStr(main_mod + "/main")
	if err != nil {
		return err
	}

	if a, ok := main_fn.(MetaApply); ok {
		l_args := data.ListFromSlice(args)
		_, err := a.MetaApply(rt.Ctx(), l_args)
		return err
	} else {
		return WrongType{"gle", main_fn, "function"}
	}
}

func RunMain(main string, args []string) {
	load(main)
	fmt.Printf("%s %#v\n", main, args)
}
