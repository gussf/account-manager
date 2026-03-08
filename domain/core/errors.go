package core

import "errors"

var ErrValidation = errors.New("validation error")
var ErrNotFound = errors.New("resource not found")
var ErrAlreadyExists = errors.New("resource already exists")
