package main

import (
	"flag"

	"github.com/Kareky/primes/config"
	"github.com/Kareky/primes/internal/bootstrap"
	database "github.com/Kareky/primes/internal/db"
)



func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	bootstrap.InitConfig(*configPath)

	defer database.Close()
	bootstrap.InitDatabase(config.Config.Database.Path)
}