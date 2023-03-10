package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/errors"
)

// Encode implements IHandler
func (h *HandlerImpl) Encode(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodPost {
		return errors.GetError(errors.InvalidRequest, fmt.Errorf("Error := %v", "Request method not allowed"))
	}
	ctx := r.Context()

	var payload forms.EncodeRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error from %v", err)
		return err
	}

	if err = json.Unmarshal(body, &payload); err != nil {
		log.Fatalf("Error from %v", err)
		return err
	}

	err = payload.Validate()
	if err != nil {
		return errors.GetError(errors.InvalidRequest, err)
	}

	encodeRes, err := h.Usecase.Encode(ctx, &payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", h.Config.Prefix, encodeRes)

	res := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Url     string `json:"url"`
	}{
		Code:    201,
		Message: "ok",
		Url:     url,
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}
