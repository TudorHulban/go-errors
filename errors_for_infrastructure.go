package goerrors

import (
	"fmt"
	"strings"
)

type ErrInfrastructure struct {
	Issue              error
	NameInfrastructure string
	Caller             string
}

func (e ErrInfrastructure) Error() string {
	var builder strings.Builder

	builder.Grow(64)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaInfrastructure)
	builder.WriteString(_space)
	builder.WriteString("Name: ")
	builder.WriteString(e.NameInfrastructure)
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

type ErrInfrastructureWParams struct {
	Issue              error
	Params             any
	NameInfrastructure string
	Caller             string
}

func (e ErrInfrastructureWParams) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaInfrastructure)
	builder.WriteString(_space)
	builder.WriteString("Name: ")
	builder.WriteString(e.NameInfrastructure)
	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)
	builder.WriteString(_space)
	builder.WriteString("Params: ")
	fmt.Fprintf(&builder, "%#v", e.Params)

	if e.Issue != nil {
		builder.WriteString(_space)
		builder.WriteString("Issue: ")
		builder.WriteString(e.Issue.Error())
	}

	return builder.String()
}
