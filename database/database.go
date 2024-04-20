package database

import (
	"api/sql/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connection open connection with database
func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnectionDatabase)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
