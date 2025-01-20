package goerrors

import "fmt"

type ErrInfrastructure struct {
	Issue              error
	NameInfrastructure string
	Caller             string
}

func (e ErrInfrastructure) Error() string {
	res := [3]string{
		sprintf("Area: %s", _AreaInfrastructure),
		sprintf("Name: %s", e.NameInfrastructure),
		fmt.Sprintf("Caller: %s", e.Caller),
	}

	if e.Issue == nil {
		return res[0] + _space + res[1] + _space + res[2]
	}

	return res[0] +
		_space +
		res[1] +
		_space +
		res[2] +
		_space +
		sprintf("Issue: %s", e.Issue.Error())
}

type ErrInfrastructureWParams struct {
	Issue              error
	Params             any
	NameInfrastructure string
	Caller             string
}

func (e ErrInfrastructureWParams) Error() string {
	res := [4]string{
		sprintf("Area: %s", _AreaInfrastructure),
		sprintf("Name: %s", e.NameInfrastructure),
		sprintf("Caller: %s", e.Caller),
		fmt.Sprintf("Params: %#v", e.Params),
	}

	if e.Issue == nil {
		return res[0] +
			_space +
			res[1] +
			_space +
			res[2] +
			_space +
			res[3]
	}

	return res[0] +
		_space +
		res[1] +
		_space +
		res[2] +
		_space +
		res[3] +
		_space +
		fmt.Sprintf("Issue: %s", e.Issue.Error())
}
