package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config holds the configuration parameters
type Config struct {
	Server struct {
		Host string
		Port string
	}
	SQLite struct {
		Path string
	}
	Option struct {
		Prefix string
	}
}

// Load returns a configuration from a JSON file
func Load(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading from file: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %v", err)
	}

	return &cfg, nil
}
