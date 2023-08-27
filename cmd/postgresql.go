package cmd

import (
	"database/sql"
	"log"
)

func newPostgresqlConection(uri string) *sql.DB {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
