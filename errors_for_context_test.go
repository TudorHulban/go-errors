package goerrors

import (
	"context"
	"errors"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrContextValueNotFound(t *testing.T) {
	ctx := context.Background()

	f := func(_ context.Context) error {
		const depth = 32

		stack := make([]uintptr, depth)
		n := runtime.Callers(2, stack[:])
		stack = stack[:n]

		return &ErrContextValueNotFound{
			stack: stack,
		}
	}

	errF := f(ctx)

	require.Error(t, errF)

	errCasted, couldCast := errF.(*ErrContextValueNotFound)
	require.True(t, couldCast)
	require.Error(t, errCasted)

	stack := errCasted.StackTrace()
	require.NotEmpty(t,
		stack,
		"stack trace should not be empty",
	)

	for _, programCounter := range stack {
		fn := runtime.FuncForPC(programCounter)
		if fn == nil {
			continue
		}

		file, line := fn.FileLine(programCounter)
		t.Logf("\t%s - %s:%d", fn.Name(), file, line)
	}
}

func TestAsIsForErrContextValueNotFound(t *testing.T) {
	errSource := ErrContextValueNotFound{
		ValueCtx: "testValue",
	}

	var errTarget ErrContextValueNotFound

	require.True(t,
		errors.As(errSource, &errTarget),
	)
	require.False(t,
		errors.Is(errSource, ErrContextValueNotFound{}),
	)

	errAnother := errors.New("some other error")
	require.False(t,
		errors.Is(errAnother, ErrContextValueNotFound{}),
	)
}
