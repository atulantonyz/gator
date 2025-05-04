package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	config_path, err := getConfigFilePath()
	content, err := os.Open(config_path)
	if err != nil {
		fmt.Println("Error opening config file")
		return Config{}, err
	}
	defer content.Close()
	decoder := json.NewDecoder(content)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.Current_user_name = userName
	return write(*cfg)

}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "home directory not found", err
	}
	config_filepath := filepath.Join(home_dir, configFileName)
	return config_filepath, nil
}

func write(cfg Config) error {
	config_path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	//Write data to config file
	file, err := os.Create(config_path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil

}
