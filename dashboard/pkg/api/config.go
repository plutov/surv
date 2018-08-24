package api

import (
	"encoding/json"
	"os"
)

// Config definition
type Config struct {
	Services []SurveyService
}

// GetConfig reads config file
func GetConfig() (*Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var c *Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
