package sdk

type Error struct {
	Human string
	Raw   error
}

func (e *Error) Error() string {
	return e.Raw.Error()
}
