package config

import (
	"os"
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
	// viper.AddConfigPath("../pkg/config/envs")
	// viper.SetConfigName("dev")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	// err = viper.ReadInConfig()

	// if err != nil {
	// 	return
	// }

	// err = viper.Unmarshal(&config)

	// TODO: Set config for permission type possibly in env
	config = Config{
		Port:       os.Getenv("AUTH_SVC_PORT"),
		DBURL:      os.Getenv("AUTH_DB_URL"),
		DBUSER:     os.Getenv("AUTH_DB_USER"),
		DBPWD:      os.Getenv("AUTH_DB_PWD"),
		APP_URL:    os.Getenv("APP_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	return config, err
}
