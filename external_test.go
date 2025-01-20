package goerrors_test

import (
	"testing"

	goerrors "github.com/TudorHulban/go-errors"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	require.Error(t,
		goerrors.ErrUnknownValue,
	)
}
