package main

import (
	"flag"

	"github.com/Kareky/primes/config"
	"github.com/Kareky/primes/internal/bootstrap"
	database "github.com/Kareky/primes/internal/db"
)



func main() {
	configPath := flag.String("config", "", "path to config file")
	dbPath := flag.String("path", "", "path of the db to seed")
	dbType := flag.String("type", "sqlite", "the type of database used for the connection")
	upperBound := flag.Int("bound", 1000000, "the number up to which primes get seeded")
	algorithm  := flag.String("algo", "eratosthenes", "algorithm to use for primes generation")
	flag.Parse()

	bootstrap.InitConfig(*configPath)

	defer database.Close()
	bootstrap.InitDatabase(config.Config.Database.Path)
	bootstrap.SeedDatabase(*dbPath, *dbType, *upperBound, *algorithm)
}