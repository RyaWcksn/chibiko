package forms

type ErrorForm struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response string `json:"response"`
}
