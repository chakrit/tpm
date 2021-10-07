package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DatabaseUrl = "postgres://0.0.0.0:5432/tpm?sslmode=disable"

func connect() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", DatabaseUrl)
}
