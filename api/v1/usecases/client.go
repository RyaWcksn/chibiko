package usecases

import (
	"context"

	"github.com/RyaWcksn/chibiko/api/v1/repositories"
	"github.com/RyaWcksn/chibiko/forms"
)

type UsecaseImpl struct {
	dbPort repositories.IDatabase
}

type IUsecase interface {
	Encode(ctx context.Context, payload *forms.EncodeRequest) (resp string, err error)
}

var _ IUsecase = (*UsecaseImpl)(nil)
