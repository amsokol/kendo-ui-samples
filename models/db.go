package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func InitDB(host, database, user, pwd string) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user, pwd, host, database)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		db = nil
	}

	return db, err
}
