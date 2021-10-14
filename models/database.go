package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DatabaseUrl = "postgres://tpm:tpm@pg:5432/tpm?sslmode=disable"

func connect() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", DatabaseUrl)
}
