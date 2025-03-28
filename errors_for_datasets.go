package goerrors

import (
	"fmt"
	"strings"
)

type ErrEntryNotFound struct {
	Key any
}

func (e ErrEntryNotFound) Error() string {
	return fmt.Sprintf(
		"entry for key '%v' not found",
		e.Key,
	)
}

type ErrDatasetEntriesNotFound struct {
	Key    any
	Caller string
}

const msgErrEntriesNotFound = "no entries found"

func (e ErrDatasetEntriesNotFound) Error() string {
	var builder strings.Builder

	builder.Grow(64)

	builder.WriteString("Area:")
	builder.WriteString(_AreaDatasets)
	builder.WriteString(_space)
	builder.WriteString("Caller:")
	builder.WriteString(e.Caller)
	builder.WriteString(_space)
	builder.WriteString("Key:")
	fmt.Fprintf(&builder, "%v", e.Key)
	builder.WriteString(" : ")
	builder.WriteString(msgErrEntriesNotFound)

	return builder.String()
}

type ErrDatasetEntryAlreadyExists struct {
	Entry  any
	Caller string
}

const msgErrEntryExists = "already exists"

func (e ErrDatasetEntryAlreadyExists) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaDatasets)
	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)
	builder.WriteString(_space)
	builder.WriteString("Entry: ")
	fmt.Fprintf(&builder, "%#v", e.Entry)
	builder.WriteString(" : ")
	builder.WriteString(msgErrEntryExists)

	return builder.String()
}

type ErrDatasetScan struct {
	Issue  error
	Caller string
}

func (e ErrDatasetScan) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaDatasets)
	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)
	builder.WriteString(_space)
	builder.WriteString("Issue scan: ")
	fmt.Fprintf(&builder, "%#v", e.Issue.Error())

	return builder.String()
}

type ErrNoUpdates struct {
	Params any

	Caller string
}

func (e ErrNoUpdates) Error() string {
	var builder strings.Builder

	builder.Grow(128)

	builder.WriteString("Area: ")
	builder.WriteString(_AreaDatasets)
	builder.WriteString(_space)
	builder.WriteString("Caller: ")
	builder.WriteString(e.Caller)
	builder.WriteString(_space)
	builder.WriteString("Params: ")
	fmt.Fprintf(&builder, "%#v", e.Params)

	return builder.String()
}
