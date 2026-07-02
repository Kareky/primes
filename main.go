package main

import (
	"flag"
	"log"

	"github.com/Kareky/primes/config"
	database "github.com/Kareky/primes/internal/db"
	_ "modernc.org/sqlite"
)

func initConfig() {
	configPath := flag.String("config", "", "path to config file")
    flag.Parse()

	log.Println("Initializing configuration...")
    cfg, err := config.Load(*configPath)
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

func main() {
	initConfig()

	defer database.Close()
	initDatabase()
}