package app

import (
	"database/sql"
)

func dbConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn) //user:password@/dbname
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
