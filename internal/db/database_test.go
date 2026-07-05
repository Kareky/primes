package db_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Kareky/primes/config"
	"github.com/Kareky/primes/internal/bootstrap"
	"github.com/Kareky/primes/internal/db"
)

func BenchmarkExists(b *testing.B) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	db.Default.Exists(999999874000003957)
}

func BenchmarkGetPrimesUpTo(b *testing.B) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	db.Default.GetPrimesUpTo(999999874000003957)
}

func BenchmarkGetAllPrimes(b *testing.B) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	db.Default.GetAllPrimes()
}