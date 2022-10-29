package db

import (
	"database/sql"

	"github.com/danilo-lopes/simple-bank/src/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.DatabaseStringConnection)
	if erro != nil {
		return nil, erro
	}

	return db, nil
}
