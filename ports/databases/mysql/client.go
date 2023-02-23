package mysql

import (
	"database/sql"

	"github.com/RyaWcksn/chibiko/api/v1/repositories"
)

type DBImpl struct {
	sql *sql.DB
}

// NewSql initialize MySQL instance
func NewSql(s *sql.DB) *DBImpl {
	return &DBImpl{
		sql: s,
	}
}

var _ repositories.IDatabase = (*DBImpl)(nil)
