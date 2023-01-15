package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const ConnectionString = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

var database *sql.DB

func Connect() {
	db, err := sql.Open("postgres", ConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("database Connected")
	database = db
}
