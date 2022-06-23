package errors

import "errors"

var (
	ErrCreatingNewRequest  = errors.New("ERROR: Error while creating new request")
	ErrUnauthorizedRequest = errors.New("ERROR: The request is unauthorized")
	ErrMarshalJSON         = errors.New("ERROR: Error while encode to JSON")
	ErrUnmarshalJSON       = errors.New("ERROR: Error while decode from JSON")
	ErrHittingGraphQL      = errors.New("ERROR: Error while hitting the GraphQL Github API")
	ErrReadingResponseBody = errors.New("ERROR: Error while reading response body")
)

var (
	ErrCreatingNewFolder = errors.New("ERROR: Error while creating new folder")
	ErrCreatingNewFile   = errors.New("ERROR: Error while creating new file")
	ErrWhileSync         = errors.New("ERROR: Error while sync or save new file")
)
