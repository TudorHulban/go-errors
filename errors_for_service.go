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
	area := sprintf(
		"Area: %s-%s",
		_AreaService,
		strings.ToUpper(e.NameService[:1])+e.NameService[1:], // ASCII speed
	)

	caller := "Caller: " + e.Caller

	if e.Issue != nil {
		return area + _space + caller + _space + "Issue: " + e.Issue.Error()
	}

	return area + _space + caller
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
	var issueDescription string

	if e.Issue != nil {
		issueDescription = _space + `Issue: ` + e.Issue.Error()
	}

	return sprintf(
		"Area: %s-%s",
		_AreaService,
		strings.ToUpper(e.NameService[:1])+e.NameService[1:], // ASCII speed
	) +
		_space +
		`Caller: ` + e.Caller +
		issueDescription +
		_space +
		fmt.Sprintf("%#v", e.Params)
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
	area := sprintf(
		"Area: %s-%s",
		areaErrServiceValidation,
		e.ServiceName,
	)

	caller := sprintf("Caller: %s", e.Caller)

	if e.Issue != nil {
		return area + _space + caller + _space + "Issue: " + e.Issue.Error()
	}

	return area + _space + caller
}
