package db

import (
	"database/sql"
	"log"
)

func DbConnect() *sql.DB {
	connStr := "user=postgres dbname=alurapg password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
