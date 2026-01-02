package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	workingDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := workingDirectory + "/" + configFileName
	return filePath, nil
}

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(c)
}

func Read() (Config, error) {
	var config Config

	filePath, err := getConfigFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, nil
	}

	return config, nil
}

func write(config *Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
