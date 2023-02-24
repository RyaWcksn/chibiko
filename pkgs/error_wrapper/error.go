package errorwrapper

import (
	"github.com/RyaWcksn/chibiko/forms"
)

// List of all errors
const (
	InternalError  = "internal error"
	InvalidRequest = "Invalid request"
)

type ErrorForm struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

func (e *ErrorForm) GetError() {
}

func getErrCode(errName string) forms.ErrorForm {
	formErr := forms.ErrorForm{}
	switch errName {
	case InternalError:
		formErr.Code = 500
		formErr.Message = InternalError
	case InvalidRequest:
		formErr.Code = 400
		formErr.Message = InvalidRequest
	}
	return formErr
}
