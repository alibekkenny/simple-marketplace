package model

import "errors"

var ErrNotFound = errors.New("not found")
var ErrDuplicate = errors.New("duplicate entry")
var ErrInternal = errors.New("internal server error")
var ErrInvalidInput = errors.New("invalid input")
var ErrInvalidCredentials = errors.New("invalid credentials")
