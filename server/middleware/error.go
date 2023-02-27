package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/chibiko/pkgs/errors"
)

type ErrHandler func(http.ResponseWriter, *http.Request) error

// serveHTTP ...
func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := errors.ErrorForm{
				Code:     500,
				Message:  "Panic",
				Response: "Error",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(xerr.Code)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := err.(*errors.ErrorForm)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(xerr.Code)
		json.NewEncoder(w).Encode(xerr)
	}

}
