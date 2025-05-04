package config

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c Config) SetUser(user string) error {
	c.Current_user_name = user
	err := write(c)
	if err != nil {
		return err
	}
	return nil

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
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	//Write data to config file
	err = os.WriteFile(config_path, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}
