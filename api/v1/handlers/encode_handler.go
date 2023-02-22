package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/RyaWcksn/chibiko/forms"
)

// Encode implements IHandler
func (h *HandlerImpl) Encode(w http.ResponseWriter, r *http.Request) error {
	_ = r.Context()

	var payload forms.EncodeRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error from ", err)
		return err
	}

	if err = json.Unmarshal(body, &payload); err != nil {
		log.Fatalf("Error from ", err)
		return err
	}

	return nil
}
