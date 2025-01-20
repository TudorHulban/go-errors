package goerrors

import "fmt"

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
	return sprintf("Area:%s", _AreaDatasets) +
		_space +
		sprintf("Caller:%s", e.Caller) +
		_space +
		fmt.Sprintf("Key:%v", e.Key) +
		" : " +
		msgErrEntriesNotFound
}

type ErrDatasetEntryAlreadyExists struct {
	Entry  any
	Caller string
}

const msgErrEntryExists = "already exists"

func (e ErrDatasetEntryAlreadyExists) Error() string {
	return sprintf("Area: %s", _AreaDatasets) +
		_space +
		sprintf("Caller: %s", e.Caller) +
		_space +
		fmt.Sprintf("Entry: %#v", e.Entry) +
		" : " +
		msgErrEntryExists
}

type ErrDatasetScan struct {
	Issue  error
	Caller string
}

func (e ErrDatasetScan) Error() string {
	return sprintf("Area: %s", _AreaDatasets) +
		_space +
		sprintf("Caller: %s", e.Caller) +
		_space +
		fmt.Sprintf("Issue scan: %#v", e.Issue.Error())
}
