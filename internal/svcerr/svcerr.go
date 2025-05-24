package svcerr

import (
	"errors"
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

func (s *SvcErr) Error() string {
	return s.serviceErr.Error() + strings.Join(s.msgs, sepStr)
}

func (s *SvcErr) Unwrap() error {
	return s.serviceErr
}

func New(baseErr BaseErr, serviceErr error, msgs ...string) error {
	return &SvcErr{
		BaseErr:    baseErr,
		serviceErr: serviceErr,
		msgs:       msgs,
	}
}
