package division_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Kareky/primes/config"
	"github.com/Kareky/primes/internal/bootstrap"
	"github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/primality/division"
)

func TestIsPrime(t *testing.T) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	maxPrime, err := db.Default.GetMaxPrime()
	if err != nil {
		t.Errorf("IsPrime(GetMaxPrime) error = %v", err)
	}
	log.Printf("Highest prime is: %d. It's pow is: %d", maxPrime, maxPrime*maxPrime)
	tests := []struct {
		name    string
		n       int
		want    bool
		wantErr bool
	}{
		// Edge cases
		{"negative", -5, false, false},
		{"zero", 0, false, false},
		{"one", 1, false, false},
		{"two", 2, true, false},
		{"three", 3, true, false},

		// Small primes
		{"five", 5, true, false},
		{"seven", 7, true, false},
		{"eleven", 11, true, false},
		{"thirteen", 13, true, false},

		// Small composites
		{"four", 4, false, false},
		{"six", 6, false, false},
		{"eight", 8, false, false},
		{"nine", 9, false, false},
		{"ten", 10, false, false},
		{"twenty five", 25, false, false},
		{"hundred", 100, false, false},

		// Larger primes
		{"101", 101, true, false},
		{"103", 103, true, false},
		{"107", 107, true, false},
		{"999983", 999983, true, false}, // known prime

		// Larger composites
		{"999966", 999966, false, false}, // even
		{"999982", 999982, false, false}, // even

		// Error case: exceeds size limit
		{"exceeds limit", maxPrime*maxPrime+1, false, true}, // adjust based on your sizeLimit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := division.IsPrime(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkIsPrime_Prime(b *testing.B) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	for b.Loop() {
		division.IsPrime(999999874000003969)
	}
}

func BenchmarkIsPrime_Composite(b *testing.B) {
	bootstrap.InitConfig("../../test.yaml")
	abs, _ := filepath.Abs(config.Config.Database.Path)
	log.Printf("Opening database with path %s", config.Config.Database.Path)
	log.Printf("Opening database at: %s", abs)
	dir := filepath.Dir(config.Config.Database.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", dir)
	}
	bootstrap.InitDatabase(config.Config.Database.Path)

	for b.Loop() {
		division.IsPrime(999966)
	}
}
