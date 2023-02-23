package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database Database `json:'database'`
	Redis    Redis    `json:'redis'`
	Prefix   string   `json:"prefix"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Database string `json:"database"`
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
