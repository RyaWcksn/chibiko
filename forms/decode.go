package forms

import (
	"fmt"

	"github.com/go-playground/validator"
)

type DecodePayload struct {
	Param string
}

func (s *DecodePayload) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("[ValidateCreateBadgeRequest] Invalid request in field [%+v], tag [%+v], value [%+v]", err.StructNamespace(), err.Tag(), err.Param())
		}
	}
	return nil
}
