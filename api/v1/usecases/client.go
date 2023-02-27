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

// NewUsecase initiate usecase layer
func NewUsecase(dbPort repositories.IDatabase) *UsecaseImpl {
	return &UsecaseImpl{
		dbPort: dbPort,
	}
}

var _ IUsecase = (*UsecaseImpl)(nil)
