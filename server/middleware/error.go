package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/errors"
)

type ErrHandler func(http.ResponseWriter, *http.Request) error

// serveHTTP ...
func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := forms.ErrorForm{
				Code:     503,
				Message:  "Error!",
				Response: "Internal Server Error",
			}
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(xerr)
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := err.(*errors.ErrorForm)
		w.WriteHeader(xerr.Code)
		json.NewEncoder(w).Encode(xerr)
	}

}
