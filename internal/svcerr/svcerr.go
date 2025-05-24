package svcerr

import (
	"errors"
	"fmt"
	"strings"
)

const sepStr = ", "

type BaseErr error

var (
	ErrInternal BaseErr = errors.New("internal")
	ErrNotFound BaseErr = errors.New("not found")
	ErrBadReq   BaseErr = errors.New("bad request")
)

type SvcErr struct {
	BaseErr    BaseErr
	serviceErr error
	msgs       []string
}

func (e *SvcErr) Error() string {
	return fmt.Sprintf("%s: %s", e.serviceErr.Error(), strings.Join(e.msgs, sepStr))
}

func (e *SvcErr) Unwrap() error {
	return e.serviceErr
}

func New(baseErr BaseErr, serviceErr error, msgs ...string) error {
	return &SvcErr{
		BaseErr:    baseErr,
		serviceErr: serviceErr,
		msgs:       msgs,
	}
}
