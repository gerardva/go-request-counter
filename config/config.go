package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	CacheHost     string `mapstructure:"REDIS_HOST"`
	CachePassword string `mapstructure:"REDIS_PASSWORD"`
	ServerPort    string `mapstructure:"PORT"`
	ClientOrigin  string `mapstructure:"CLIENT_ORIGIN"`
	CountKey      string `mapstructure:"COUNT_KEY"`
	TotalCountKey string `mapstructure:"TOTAL_COUNT_KEY"`
	Hostname      string
}

var config *Config

func Init() {
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$GOPATH/src/github.com/gerardva/go-api/")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.Unmarshal(&config)
	config.Hostname = viper.GetString("HOSTNAME")
}

func GetConfig() *Config {
	return config
}
