package data

import (
	"fmt"
	"strings"
)

func str1(val any) string {
	if val == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", val)
}

func Str(values ...any) string {
	var strs []string
	for _, v := range values {
        strs = append(strs, str1(v))
	}
	return strings.Join(strs, " ")
}
