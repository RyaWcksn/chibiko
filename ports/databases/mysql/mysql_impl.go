package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RyaWcksn/chibiko/entities"
	"github.com/RyaWcksn/chibiko/pkgs/errors"
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

func (s *DBImpl) Get(ctx context.Context, entity *entities.GetDatabase) (url string, err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	fmt.Println("Masuk sini ga nih")
	var urlResp string

	tx, err := s.sql.Begin()
	if err != nil {
		log.Printf("Error := %v", err)
		return "", errors.GetError(errors.InternalServer, err)
	}
	err = s.sql.QueryRowContext(ctxTimeout, `UPDATE urls set count = count + 1 where id = ?`, entity.Id).Err()
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Printf("Error := %v", err)
			return "", errors.GetError(errors.UnavailableService, err)
		}
		log.Printf("Error := %v", err)
		return "", errors.GetError(errors.UnavailableService, err)
	}
	err = s.sql.QueryRowContext(ctxTimeout, `SELECT url FROM urls where id = ?`, entity.Id).Scan(&urlResp)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Printf("Error := %v", err)
			return "", errors.GetError(errors.InvalidRequest, err)
		}
		log.Printf("Error := %v", err)
		return "", errors.GetError(errors.UnavailableService, err)
	}

	return urlResp, nil
}
