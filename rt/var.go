package rt

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
