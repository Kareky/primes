package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

var Config Settings

type Settings struct {
	Database 	Database `yaml:"database"`
}

type Database struct {
    Enabled    	bool   `yaml:"enabled"`
    Type       	string `yaml:"type"`
    Path	  	string `yaml:"path"`
    Name       	string `yaml:"name"`
    User       	string `yaml:"user"`
    Pass       	string `yaml:"pass"`
    UpperBound 	int    `yaml:"upper-bound"`
}

func Default() Settings {
    return Settings{
		Database{
			Enabled: 	true,
			Type:		"sqlite",
			Path:		"./data/primes.db",
			Name:		"primes",
			User:		"",
			Pass:		"",
			UpperBound:	1000000000000,
		},
    }
}


// Load reads a YAML file from the given path and unmarshals it into Settings.
// If the file doesn't exist, it returns the default config.
func Load(path string) (Settings, error) {
    cfg := Default()

    if path == "" {
        return cfg, nil // no file specified, use defaults
    }

    data, err := os.ReadFile(path)
    if err != nil {
        if os.IsNotExist(err) {
            return cfg, nil // file doesn't exist, use defaults
        }
        return cfg, fmt.Errorf("failed to read config file %s: %w", path, err)
    }

    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        return cfg, fmt.Errorf("failed to parse config file %s: %w", path, err)
    }

    return cfg, nil
}