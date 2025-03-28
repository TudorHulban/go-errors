package goerrors

import (
	"strings"
)

type ErrCaching struct {
	Issue       error
	NameCaching string // ex. "LRU", ASCII only.
	Caller      string
}

func (e ErrCaching) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaCache)
	builder.WriteByte('-')

	if len(e.NameCaching) > 0 {
		builder.WriteString(strings.ToUpper(e.NameCaching[:1]))
		if len(e.NameCaching) > 1 {
			builder.WriteString(e.NameCaching[1:])
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

func (e ErrCaching) Unwrap() error {
	return e.Issue
}
