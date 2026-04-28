package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	ConfigPath      string `env:"CONFIG_PATH" envDefault:"./config.json"`
	PrivateKeyPath  string
	RecPath         string
	ServerAddress   string
	SessionFilePath string
	StoreDirPath    string
}

type jsonConfig struct {
	PrivateKeyPath  string `json:"privateKeyPath"`
	PublicKeyPath   string `json:"publicKeyPath"`
	ServerAddress   string `json:"serverAddress"`
	SessionFilePath string `json:"sessionFilePath"`
	StoreDirPath    string `json:"storeDirPath"`
}

func InitConfig() (*Config, error) {
	envValues := initEnv()

	return assignJSONConfig(envValues)
}

func initEnv() *Config {
	conf := Config{}

	env.Parse(&conf)

	return &conf
}

func assignJSONConfig(config *Config) (*Config, error) {
	if config.ConfigPath != "" {
		_, err := os.Stat(config.ConfigPath)

		if err == nil || !errors.Is(err, os.ErrNotExist) {
			file, err := os.OpenFile(config.ConfigPath, os.O_RDONLY, 0x666)
			if err == nil {
				defer file.Close()
				var jsonConfig jsonConfig

				if err := json.NewDecoder(file).Decode(&jsonConfig); err == nil {
					config.PrivateKeyPath = jsonConfig.PrivateKeyPath
					config.RecPath = jsonConfig.PublicKeyPath
					config.SessionFilePath = jsonConfig.SessionFilePath
					config.ServerAddress = jsonConfig.ServerAddress
					config.StoreDirPath = jsonConfig.StoreDirPath
				}
			}
		}
	}

	return config, nil
}
