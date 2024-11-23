package helpers

import (
	"runtime"
)

type Error struct {
	Code                   int           `json:"code"`
	ErrorMessageForClient  string        `json:"error_message_for_client"`
	ErrorDetailForEngineer string        `json:"error_detail_for_engineer"`
	ErrorLocation          ErrorLocation `json:"error_location"`
}

type ErrorLocation struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

func (e *Error) Error() string {
	return e.ErrorMessageForClient
}

func NewError(code int, errorForClient, errorForEngineer string) *Error {

	_, file, line, _ := runtime.Caller(1)

	err := &Error{
		Code:                   code,
		ErrorMessageForClient:  errorForClient,
		ErrorDetailForEngineer: errorForEngineer,
		ErrorLocation: ErrorLocation{
			File: file,
			Line: line,
		},
	}

	return err
}
