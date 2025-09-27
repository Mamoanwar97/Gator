package config

import (
	"encoding/json"
	"os"
)

func LoadConfig() (Config, error) {
	configPath := GetConfigPath()
	config, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	var configData Config
	err = json.Unmarshal(config, &configData)
	if err != nil {
		return Config{}, err
	}
	return configData, nil
}

func SetUser(config *Config, userName string) error {
	config.CurrentUserName = userName
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}
	os.WriteFile(GetConfigPath(), jsonData, 0644)
	return nil
}
