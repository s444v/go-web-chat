package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func DBinit() error {
	var err error
	connStr := "user=postgres password=1 dbname=web-chat sslmode=disable"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}
	return err
}
