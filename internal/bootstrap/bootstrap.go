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

func InitConfig(configPath string) {
	log.Println("Initializing configuration...")
    cfg, err := config.Load(configPath)
    if err != nil {
        log.Fatal(err)
    }

	config.Config = cfg

	log.Println("Configuration initialized")
}

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