package repositories

import (
	"context"

	"github.com/RyaWcksn/chibiko/entities"
)

//go:generate mockgen -source db_repository.go -destination db_repository_mock.go -package repositories
type IDatabase interface {
	Save(ctx context.Context, entity *entities.SaveDatabase) (id int64, err error)
	Get(ctx context.Context, entity *entities.GetDatabase) (url string, err error)
}
