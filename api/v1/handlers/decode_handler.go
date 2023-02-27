package handlers

import (
	"fmt"
	"net/http"

	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/errors"
)

func (h *HandlerImpl) Decode(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodGet {
		return errors.GetError(errors.InvalidRequest, fmt.Errorf("Error := %v", "Request method not allowed"))
	}

	code := r.URL.Path[len("/"):]

	payload := forms.DecodePayload{
		Param: code,
	}
	fmt.Println(code)

	err := payload.Validate()
	if err != nil {
		return errors.GetError(errors.InvalidRequest, err)
	}

	resp, err := h.Usecase.Decode(r.Context(), &payload)
	if err != nil {
		return err
	}
	fmt.Println(resp)

	http.Redirect(w, r, resp, http.StatusMovedPermanently)
	return nil
}
