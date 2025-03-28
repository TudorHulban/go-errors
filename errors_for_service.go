package goerrors

import (
	"fmt"
	"strings"
)

type ErrService struct {
	Issue       error
	NameService string // ex. "Ticket", ASCII only.
	Caller      string
}

func (e ErrService) Error() string {
	var builder strings.Builder

	builder.Grow(64)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaService)
	builder.WriteByte('-')

	if len(e.NameService) > 0 {
		builder.WriteString(strings.ToUpper(e.NameService[:1]))
		if len(e.NameService) > 1 {
			builder.WriteString(e.NameService[1:])
		}
	}

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

func (e ErrService) Unwrap() error {
	return e.Issue
}

type ErrServiceWithParams struct {
	Params      any
	Issue       error
	NameService string // ex. "Ticket", ASCII only.
	Caller      string
}

func (e ErrServiceWithParams) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaService)
	builder.WriteByte('-')
	if len(e.NameService) > 0 {
		first := strings.ToUpper(e.NameService[:1])
		builder.WriteString(first)
		if len(e.NameService) > 1 {
			builder.WriteString(e.NameService[1:])
		}
	}

	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)

	if e.Issue != nil {
		builder.WriteString(_space)
		builder.WriteString("Issue: ")
		builder.WriteString(e.Issue.Error())
	}

	builder.WriteString(_space)
	fmt.Fprintf(&builder, "%#v", e.Params)

	return builder.String()
}

func (e ErrServiceWithParams) Unwrap() error {
	return e.Issue
}

type ErrServiceValidationNoUpdates struct {
	Caller string
}

const areaErrServiceValidationNoUpdates = "Service-Validation-NoUpdates"

func (e ErrServiceValidationNoUpdates) Error() string {
	return sprintf("Area: %s", areaErrServiceValidationNoUpdates) +
		_space +
		"Caller: " + e.Caller
}

type ErrServiceValidation struct {
	Issue       error
	ServiceName string // ex. "Ticket", ASCII only.
	Caller      string
}

const areaErrServiceValidation = "Service-Validation"

func (e ErrServiceValidation) Error() string {
	var builder strings.Builder

	builder.Grow(64)

	builder.WriteString("Area: ")
	builder.WriteString(areaErrServiceValidation)
	builder.WriteByte('-')
	builder.WriteString(e.ServiceName)

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
