package errors

import "errors"

var (
	IncorrectEmailOrPassword = errors.New("incorrect email or password")
	NotAuthenticated         = errors.New("not authenticated")
)
