package postgres

import (
	"database/sql"
	"udv-test-task-go/internal/usecase"
)

type pgGetListDB struct {
	Conn *sql.DB
}

func NewGetListDB(Conn *sql.DB) usecase.GetListDBInterface {
	return &pgGetListDB{
		Conn: Conn,
	}
}

func (p *pgGetListDB) GetListDatabases() ([]string, error) {
	rows, err := p.Conn.Query("SELECT datname FROM pg_database")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []string

	for rows.Next() {
		var dbname string
		err := rows.Scan(&dbname)
		if err != nil {
			return nil, err
		}
		databases = append(databases, dbname)
	}
	return databases, nil
}
