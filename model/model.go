package model

import "errors"

var (
	ErrPersonCanNotBeNil     = errors.New("person can't be nil")
	ErrIDPersonDoesNotExists = errors.New("the ID doesn't exists")
)
