package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/Kareky/primes/config"
	database "github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/internal/seeder"
	_ "modernc.org/sqlite"
)

func initConfig(configPath string) {
	log.Println("Initializing configuration...")
    cfg, err := config.Load(configPath)
    if err != nil {
        log.Fatal(err)
    }

	config.Config = cfg

	log.Println("Configuration initialized")
}

func initDatabase() {
	log.Printf("Initializing database at path %s...", config.Config.Database.Path)
	err := database.Initialize(config.Config.Database.Path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized")
}

func seedDatabase(dbPath string, dbType string, upperBound int, algorithm string) {
	var db *database.DB
	if dbPath != "" {
		sqlDB, err := sql.Open(dbType, dbPath)
		if err != nil {
			log.Fatal(err)
		}

		db, err = database.NewDB(sqlDB)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db = database.Default
	}

	seeder.NewSeed(db).PopulatePrimesSeed(upperBound, algorithm)
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	dbPath := flag.String("dbPath", "", "path of the db to seed")
	dbType := flag.String("dbType", "sqlite", "the type of database used for the connection")
	upperBound := flag.Int("uppperBound", 1000000, "the number up to which primes get seeded")
	algorithm  := flag.String("algorithm", "eratosthenes", "algorithm to use for primes generation")
	flag.Parse()

	initConfig(*configPath)

	defer database.Close()
	initDatabase()
	seedDatabase(*dbPath, *dbType, *upperBound, *algorithm)
}