package goerrors

import (
	"errors"
	"fmt"
)

var ErrUnknownValue = errors.New("unknown value")

type ErrValidation struct {
	Issue  error
	Caller string
}

func (e ErrValidation) Error() string {
	area := "Area: " + _AreaValidation
	caller := "Caller: " + e.Caller

	if e.Issue != nil {
		return area + _space + caller + _space + "Issue: " + e.Issue.Error()
	}

	return area + _space + caller
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
