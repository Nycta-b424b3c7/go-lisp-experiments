package main

type Memo[T any] struct {
	hasValue bool
	getValue func() T
	value    T
}

func MakeMemo[T any](f func() T) Memo[T] {
	m := Memo[T]{}
	m.getValue = f
	return m
}

func (m *Memo[T]) Get() T {
	if !m.hasValue {
		m.value = m.getValue()
	}
	return m.value
}
