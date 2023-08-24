package config

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	CertificateFilenames []string
}

func LoadConfig() (*Config, error) {
	// Read the configuration file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	// Parse the configuration data into the Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// Return the populated Config struct and any error encountered
	return &config, nil
}

