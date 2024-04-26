package data

import (
	. "gle/etc"
	"strings"
)

type seq struct {
	val  any
	rest *seq
}

func seqEq(a, b *seq) bool {
	for {
		if a == nil && b == nil {
			return true
		}
		if !IsEq(a.val, b.val) {
			return false
		}
		a = a.rest
		b = b.rest
	}
}

type List struct {
	count int
	seq   *seq
}

var EMPTY_LIST = List{0, nil}

func ListFromSlice[T any](slice []T) List {
	count := len(slice)
	var s *seq = nil
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		s = &seq{v, s}
	}
	return List{count, s}
}

func (l List) Count() int {
	return l.count
}

func (l List) First() any {
	if l.seq == nil {
		return nil
	}
	return l.seq.val
}

func (l List) ToSlice() []any {
	slice := make([]any, l.count)
	s := l.seq
	for i := 0; i < l.count; i++ {
		slice[i] = s.val
		s = s.rest
	}
	return slice
}

func (l List) Rest() List {
	if l.seq == nil {
		return EMPTY_LIST
	}

	return List{l.count - 1, l.seq.rest}
}

func (l List) Cons(v any) List {
	return List{
		l.count + 1,
		&seq{v, l.seq},
	}
}

func (l List) Eq(other any) bool {
	o, ok := other.(List)
	if !ok {
		return false
	}
	if o.count != o.count {
		return false
	}
	return seqEq(l.seq, o.seq)
}

func (l List) String() string {
	var b strings.Builder
	b.WriteRune('(')
	s := l.seq
	for s != nil {
		b.WriteString(Str(s.val))
		s = s.rest
		if s != nil {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(')')
	return b.String()
}

func Reduce[T any](f func(T, any) T, init T, l List) T {
	res := init
	s := l.seq
	for s != nil {
		res = f(res, s.val)
		s = s.rest
	}
	return res
}
