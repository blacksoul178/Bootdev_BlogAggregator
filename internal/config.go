package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	jsonLocation := fmt.Sprintf("%s.gatorconfig.json", homeDir)
	file, err := os.Open(jsonLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config, nil
}

func (Config) SetUser(config *Config) error {
	jsonLocation
}
