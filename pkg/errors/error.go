package errors

import (
	"fmt"
)

type Status int

const (
	StatusSuccess Status = iota
	StatusInternal
	StatusBadRequest
)

type Error struct {
	err    error
	msg    string
	status Status
}

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.msg
}

func News(msg string, args ...interface{}) error {
	return &Error{
		err:    fmt.Errorf(msg, args...),
		status: StatusInternal,
	}
}

func BadRequest(msg string, args ...interface{}) error {
	return &Error{
		err:    nil,
		msg:    fmt.Sprintf(msg, args...),
		status: StatusBadRequest,
	}
}
