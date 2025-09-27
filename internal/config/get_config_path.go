package config

import (
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, "bootDev", configFileName)
}
