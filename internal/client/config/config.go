package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/caarlos0/env"
	"github.com/spf13/pflag"
)

type Config struct {
	ConfigPath      string `env:"CONFIG_PATH"`
	PrivateKeyPath  string
	PublicKeyPath   string
	ServerAddress   string
	SessionFilePath string
}

type jsonConfig struct {
	PrivateKeyPath  string `json:"privateKeyPath"`
	PublicKeyPath   string `json:"publicKeyPath"`
	ServerAddress   string `json:"serverAddress"`
	SessionFilePath string `json:"sessionFilePath"`
}

func InitConfig() (*Config, error) {
	envValues := initEnv()
	flagValues := initFlags()

	if envValues.ConfigPath == "" {
		envValues.ConfigPath = flagValues.ConfigPath
	}

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
					config.PublicKeyPath = jsonConfig.PublicKeyPath
					config.SessionFilePath = jsonConfig.SessionFilePath
					config.ServerAddress = jsonConfig.ServerAddress
				}
			}

			return nil, err
		}
	}

	return config, nil
}

func initFlags() *Config {
	flagValues := Config{}

	pflag.StringVarP(&flagValues.ConfigPath, "conf", "c", "./config.json", "Path to config file in json format")
	pflag.Parse()

	return &flagValues
}
