package goerrors

type ErrContextValueNotFound struct {
	ValueCtx string

	stack []uintptr
}

func (e ErrContextValueNotFound) Error() string {
	return sprintf(
		"value %s not found in context",
		e.ValueCtx,
	)
}

func (e ErrContextValueNotFound) Unwrap() error {
	return nil
}

func (e *ErrContextValueNotFound) StackTrace() []uintptr {
	if e == nil {
		return nil
	}

	return e.stack
}
