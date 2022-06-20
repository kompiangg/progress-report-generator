package error

import "errors"

var (
	ErrUnauthorizedRequest = errors.New("[Error] The request is unauthorized")
)
