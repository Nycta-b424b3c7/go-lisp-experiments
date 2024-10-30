package main

import (
	"slices"
	"strings"
)

func seqEq(a, b Seq) bool {
	for {
		if a == nil && b == nil {
			return true
		}
		if !IsEq(a.First(), b.First()) {
			return false
		}
		a = a.Rest()
		b = b.Rest()
	}
}

type List struct {
	count int
	seq   Seq
}

var EMPTY_LIST = List{0, nil}

func ListFromSlice[T any](slice []T) List {
	count := len(slice)
	var s Seq = nil
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		s = &Cons{v, s}
	}
	return List{count, s}
}

func (l List) Count() int {
	return l.count
}

func (l List) GetAsSeq() Seq {
	return l.seq
}

func (l List) IsEmpty() bool {
	return l.Count() == 0
}

func (l List) First() any {
	if l.seq == nil {
		return nil
	}
	return l.seq.First()
}

func (l List) Rest() List {
	if l.seq == nil {
		return EMPTY_LIST
	}

	return List{l.count - 1, Rest(l.seq)}
}

func (l List) ToSlice() []any {
	slice := make([]any, l.count)
	s := l.seq
	for i := 0; i < l.count; i++ {
		slice[i] = First(s)
		s = Rest(s)
	}
	return slice
}

func (l List) Cons(v any) List {
	return List{
		l.count + 1,
		&Cons{v, l.seq},
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
		b.WriteString(str(First(s)))
		s = Rest(s)
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
		res = f(res, First(s))
		s = Rest(s)
	}
	return res
}

func ReduceErr[T any](f func(T, any) (T, error), init T, l List) (T, error) {
	var err error
	res := init
	s := l.seq
	for s != nil {
		res, err = f(res, First(s))
		if err != nil {
			return init, err
		}
		s = Rest(s)
	}
	return res, nil
}

func Concat(ls ...List) List {
	if len(ls) == 0 {
		return EMPTY_LIST
	}

	if len(ls) == 1 {
		return ls[0]
	}

	totalLen := 0
	for _, l := range ls {
		totalLen += l.Count()
	}

	items := make([]any, 0, totalLen)
	for _, l := range ls {
		items = slices.Concat(items, l.ToSlice())
	}

	return ListFromSlice(items)
}
