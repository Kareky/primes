package main

import (
	"log"
	_ "modernc.org/sqlite"
	database "github.com/Kareky/primes/internal/db"
)

func main() {
	err := database.Initialize("")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}