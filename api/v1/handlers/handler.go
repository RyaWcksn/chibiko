package handlers

import (
	"net/http"
)

type HandlerImpl struct {
}

type IHandler interface {
	Encode(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*HandlerImpl)(nil)
