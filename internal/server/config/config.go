package config

import (
	"github.com/caarlos0/env"
	"github.com/spf13/pflag"
)

type Config struct {
	Address       string `env:"RUN_ADDRESS"`
	DatabaseURI   string `env:"DATABASE_URI"`
	JWTSecret     string `env:"JWT_SECRET"`
	JWTExpMinutes uint   `env:"JWT_EXP_MINUTES"`
}

func InitConfig() *Config {
	envValues := initEnv()
	flagValues := initFlags()

	if envValues.Address == "" {
		envValues.Address = flagValues.Address
	}

	if envValues.DatabaseURI == "" {
		envValues.DatabaseURI = flagValues.DatabaseURI
	}

	if envValues.JWTSecret == "" {
		envValues.JWTSecret = flagValues.JWTSecret
	}

	if envValues.JWTExpMinutes == 0 {
		envValues.JWTExpMinutes = flagValues.JWTExpMinutes
	}

	return envValues
}

func initEnv() *Config {
	agentConfig := Config{}

	env.Parse(&agentConfig)

	return &agentConfig
}

func initFlags() *Config {
	flagValues := Config{}

	pflag.StringVarP(&flagValues.Address, "addr", "a", "localhost:3000", "Address host:port")
	pflag.StringVarP(&flagValues.DatabaseURI, "dbaddr", "d", "postgresql://localhost/postgres", "PG URI")
	pflag.StringVarP(&flagValues.JWTSecret, "jwtsec", "j", "DEFAULT_SECRET", "JWT Secret") // it is not right, remove default arg
	pflag.UintVarP(&flagValues.JWTExpMinutes, "jwtexp", "s", 180, "JWT lifetime in minutes")
	pflag.Parse()

	return &flagValues
}
