package main

import (
	"flag"
	"log"

	"github.com/Kareky/primes/config"
	database "github.com/Kareky/primes/internal/db"
	_ "modernc.org/sqlite"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
    flag.Parse()

    cfg, err := config.Load(*configPath)
    if err != nil {
        log.Fatal(err)
    }

	config.Config = cfg

	err = database.Initialize(cfg.Database.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}