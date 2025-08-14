package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func DBinit() error {
	var err error
	connStr := "host=91.149.255.224 port=5432 user=s444v password=- dbname=web_chat sslmode=disable"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}
	return err
}
