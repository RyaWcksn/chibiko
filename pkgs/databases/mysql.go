package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RyaWcksn/chibiko/configs"
	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	Config configs.Database
}

// NewMysqlConnection initialize MySQL connection
func NewMysqlConnection(c configs.Database) *Conn {
	return &Conn{
		Config: c,
	}
}

type IMysql interface {
	DBConnect() *sql.DB
}

var _ IMysql = (*Conn)(nil)

// DBConnect implements IMysql
func (mysql *Conn) DBConnect() *sql.DB {
	address := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		mysql.Config.Username,
		mysql.Config.Password,
		mysql.Config.Host,
		mysql.Config.Port,
		mysql.Config.Database,
	)
	dbConn, err := sql.Open("mysql", address)

	if err != nil {
		log.Printf("Error := %v", err)
		return nil
	}

	err = dbConn.Ping()
	if err != nil {
		log.Printf("Error := %v", err)
		return nil
	}

	_, err = dbConn.Exec(`
CREATE TABLE IF NOT EXISTS urls (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  url VARCHAR(255),
  count INT
);`)
	if err != nil {
		log.Printf("error creating table: %+v\n", err)
	}

	return dbConn

}
