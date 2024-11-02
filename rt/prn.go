package rt

import (
    . "gle/data"
)

func Pr(xs ...any) {
	print(Str(xs...))
}

func Prn(xs ...any) {
	println(Str(xs...))
}
