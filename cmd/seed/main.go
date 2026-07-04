package main

import (
	"flag"

	"github.com/Kareky/primes/internal/bootstrap"
	database "github.com/Kareky/primes/internal/db"
)



func main() {
	dbPath := flag.String("path", "", "path of the db to seed")
	dbType := flag.String("type", "sqlite", "the type of database used for the connection")
	upperBound := flag.Int("bound", 1000000, "the number up to which primes get seeded")
	algorithm  := flag.String("algo", "eratosthenes", "algorithm to use for primes generation")
	flag.Parse()

	defer database.Close()
	bootstrap.SeedDatabase(*dbPath, *dbType, *upperBound, *algorithm)
}