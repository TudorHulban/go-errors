package goerrors

import (
	"errors"
	"fmt"
	"strings"
)

var ErrUnknownValue = errors.New("unknown value")

type ErrValidation struct {
	Issue  error
	Caller string
}

func (e ErrValidation) Error() string {
	var builder strings.Builder

	builder.Grow(64)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaValidation)
	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)

	if e.Issue != nil {
		builder.WriteString(_space)
		builder.WriteString("Issue: ")
		builder.WriteString(e.Issue.Error())
	}

	return builder.String()
}

type ErrNilInput struct {
	InputName string
}

func (e ErrNilInput) Error() string {
	return "nil Input, name: " + e.InputName

}

type ErrInvalidInput struct {
	Issue      error
	InputValue any
	InputName  string
	Caller     string
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf(
		"caller: '%s' - invalid input, name: %s, value: %v, issue: %v",
		e.Caller,
		e.InputName,
		e.InputValue,
		e.Issue,
	)
}

type ErrNegativeInput struct {
	InputName string
}

func (e ErrNegativeInput) Error() string {
	return "negative input name: " + e.InputName
}

type ErrZeroInput struct {
	InputName string
}

func (e ErrZeroInput) Error() string {
	return "zero Input Name: " + e.InputName
}

type ErrNoMatchForValue struct {
	Value     any
	ValueName string
}

func (e ErrNoMatchForValue) Error() string {
	return fmt.Sprintf(
		"value name: '%s' - value: %v",
		e.ValueName,
		e.Value,
	)
}
