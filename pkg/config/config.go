package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"AUTH_SVC_PORT"`
	DBURL      string `mapstructure:"AUTH_DB_URL"`
	DBUSER     string `mapstructure:"AUTH_DB_USER"`
	DBPWD      string `mapstructure:"AUTH_DB_PWD"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	APP_NAME   string `mapstructure:"APP_NAME"`
	APP_URL    string `mapstructure:"APP_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
// 	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

// 	TODO: Set config for permission type possibly in env

	return config, err
}
