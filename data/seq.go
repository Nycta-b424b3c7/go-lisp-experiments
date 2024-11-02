package data

type Seq interface {
	Seq() Seq
	IsEmpty() bool
	First() any
	Rest() Seq
}

type Cons struct {
	val  any
	rest Seq
}

func (c Cons) Seq() Seq {
	return c
}

func (c Cons) IsEmpty() bool {
	return false
}

func (c Cons) First() any {
	return c.val
}

func (c Cons) Rest() Seq {
	return c.rest
}

func SliceToSeq(slice []any) Seq {
	if len(slice) == 0 {
		return nil
	}

	return Cons{slice[0], SliceToSeq(slice[1:])}
}

type LazySeq struct {
	val  Memo[any]
	rest Seq
}

func LazyCons(f func() any, rest Seq) Seq {
	return &LazySeq{MakeMemo(f), rest}
}

func (s LazySeq) Seq() Seq {
	return s
}

func (s LazySeq) IsEmpty() bool {
	return false
}

func (s LazySeq) First() any {
	return s.val.Get()
}

func (s LazySeq) Rest() Seq {
	return s.rest
}

func IsEmpty(seq Seq) bool {
	if seq != nil {
		return seq.IsEmpty()
	}

	return true
}

func First(seq Seq) any {
	if seq != nil {
		return seq.First()
	}

	return nil
}

func Rest(seq Seq) Seq {
	if seq != nil {
		return seq.Rest()
	}

	return nil
}

func seq(s Seq) Seq {
	if s != nil {
		s.Seq()
	}
	return nil
}

func SeqIntoSlice(s Seq) []any {
	into := []any{}
	seq := seq(s)
	for seq != nil {
		into = append(into, First(seq))
		seq = Rest(seq)
	}
	return into
}
