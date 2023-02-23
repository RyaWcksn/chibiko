package mysql

import (
	"context"
	"log"
	"time"

	"github.com/RyaWcksn/chibiko/entities"
)

// Save implements repositories.IDatabase
func (s *DBImpl) Save(ctx context.Context, entity *entities.SaveDatabase) (id int64, err error) {

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	tx, err := s.sql.Begin()
	if err != nil {
		log.Printf("Error := %v", err)
		return 0, err
	}

	queryString := "insert into urls(url, count) values (?, ?)"
	stmt, err := s.sql.Prepare(queryString)
	if err != nil {
		log.Printf("Error := %v", err)
		return 0, err
	}
	sqlRes, err := stmt.ExecContext(ctxTimeout, entity.Url, 0)
	if err != nil {
		tx.Rollback()
		log.Printf("Error := %v", err)
		return 0, err
	}

	res, err := sqlRes.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("Error on transaction := %v", err)
		return 0, err
	}

	return res, nil
}
