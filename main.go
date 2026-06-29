package main

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	database "primes/internal/db"
)

func main() {
	db, err := sql.Open("sqlite", "primes.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db, err = database.NewDB(db)
	if err != nil {
		log.Fatal(err)
	}
}