package goerrors

type ErrContextValueNotFound struct {
	ValueCtx string
}

func (e ErrContextValueNotFound) Error() string {
	return sprintf(
		"value %s not found in context",
		e.ValueCtx,
	)
}
