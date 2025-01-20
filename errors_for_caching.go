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
	area := sprintf(
		"Area: %s-%s",
		_AreaCache,
		strings.ToUpper(e.NameCaching[:1])+e.NameCaching[1:], // ASCII speed
	)

	caller := "Caller: " + e.Caller

	if e.Issue != nil {
		return area + _space + caller + _space +
			sprintf(
				"Issue: %s",
				e.Issue.Error(),
			)
	}

	return area + _space + caller
}

func (e ErrCaching) Unwrap() error {
	return e.Issue
}
