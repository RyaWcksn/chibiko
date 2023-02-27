package errors

import "fmt"

// List of all errors
const (
	InternalServer     = "internal error"
	InvalidRequest     = "Invalid request"
	UnavailableService = "Unavailable Service"
	Unauthorized       = "Unauthorized"
)

const (
	UnavailableServiceCode = 503
	InternalServerCode     = 500
	InvalidRequestCode     = 400
	UnauthorizedCode       = 401
)

type IError interface {
	Error() string
	GetHTTPCode() int
}

type ErrorForm struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

func (o ErrorForm) Error() string {
	return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message, o.Response)
}

func (o ErrorForm) GetHTTPCode() int {
	return o.Code
}

// getErrorCode return http error code base on error message.
func getErrorCode(errMsg string) int {
	var val int
	switch errMsg {
	case InternalServer:
		val = InternalServerCode
	case UnavailableService:
		val = UnavailableServiceCode
	case InvalidRequest:
		val = InvalidRequestCode
	case Unauthorized:
		val = UnauthorizedCode
	default:
		val = InternalServerCode
	}
	return val
}

// GetError code and message then return.
func GetError(errMessage string, errActual error) *ErrorForm {
	return &ErrorForm{
		Code:     getErrorCode(errMessage),
		Message:  errMessage,
		Response: errActual.Error(),
	}
}
