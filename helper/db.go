package helper

import (
	"go-docker/config"
	_ "github.com/lib/pq" // here
	"github.com/jmoiron/sqlx"
)

func GetDBClient() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", config.GetEnv().DBUri)
}
