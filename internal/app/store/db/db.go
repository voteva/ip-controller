package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(dbURL string) error {
	con, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	if err := con.Ping(); err != nil {
		return err
	}
	DB = con
	return nil
}

func Close() {
	DB.Close()
}
