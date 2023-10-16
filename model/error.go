package model

import "errors"

var (
	ErrRequestNotValid       = errors.New("request data is not valid")
	ErrNotFound              = errors.New("data not found")
	ErrResourceAlreadyExists = errors.New("resource already exists")
)
