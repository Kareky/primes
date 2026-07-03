// In this packages are contained all function that can be run at the start of the program
package bootstrap

import (
	"database/sql"
	"log"

	"github.com/Kareky/primes/config"
	database "github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/internal/seeder"
	"github.com/Kareky/primes/primality/division"
	_ "modernc.org/sqlite"
)

// InitConfig initializes the configuration by loading it from the specified path and setting it in the global config variable.
func InitConfig(configPath string) {
	log.Println("Initializing configuration...")
    cfg, err := config.Load(configPath)
    if err != nil {
        log.Fatal(err)
    }

	config.Config = cfg

	log.Println("Configuration initialized")
}

// InitDatabase initializes the database at the specified path. It creates a new database connection and updates the size limit
// for the division primality test based on the maximum prime number stored in the database.
func InitDatabase(databasePath string) {
	log.Printf("Initializing database at path %s...", databasePath)
	err := database.Initialize(databasePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := division.UpdateSizeLimit(); err != nil {
        log.Fatal(err)
    }

	log.Println("Database initialized")
}

// SeedDatabase seeds the database with prime numbers up to the specified upper bound using the specified algorithm.
// If a database path is provided, it creates a new database connection; otherwise, it uses the default database connection.
func SeedDatabase(dbPath string, dbType string, upperBound int, algorithm string) {
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
	if err := division.UpdateSizeLimit(); err != nil {
        log.Fatal(err)
    }
}