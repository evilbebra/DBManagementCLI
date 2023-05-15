package database

import (
	"database/sql"
	"fmt"
	"udv-test-task-go/config"
)

type DB struct {
	Conn *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(cfg config.Config) (*DB, error) {
	dbSource := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
	)
	d, err := sql.Open(cfg.Driver, dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.Conn = d
	return dbConn, err
}
