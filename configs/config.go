package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
		Password string `json:"password"`
		Username string `json:"username"`
	} `json:'database'`
	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:'redis'`
	Prefix string `json:"prefix"`
}

// ReadFromFile json env.
func ReadFromFile(path string) (conf *Config, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
