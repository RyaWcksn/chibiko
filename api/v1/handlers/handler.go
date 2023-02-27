package handlers

import (
	"net/http"

	"github.com/RyaWcksn/chibiko/api/v1/usecases"
	"github.com/RyaWcksn/chibiko/configs"
)

type HandlerImpl struct {
	Config  configs.Config
	Usecase usecases.IUsecase
}

// NewHandler initiate handler impl.
func NewHandler(cfg configs.Config, usecase usecases.IUsecase) *HandlerImpl {
	return &HandlerImpl{
		Config:  cfg,
		Usecase: usecase,
	}
}

type IHandler interface {
	Encode(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*HandlerImpl)(nil)
