package config

import (
	"os"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBURL      string `mapstructure:"DB_URL"`
	DBUSER     string `mapstructure:"DB_USER"`
	DBPWD      string `mapstructure:"DB_PWD"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	APP_NAME   string `mapstructure:"JWT_SECRET"`
	APP_URL    string `mapstructure:"JWT_SECRET"`
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

	config = Config{
		Port:    os.Getenv("Port"),
		DBURL:   os.Getenv("DB_URL"),
		DBUSER:  os.Getenv("DB_USER"),
		DBPWD:   os.Getenv("DB_PWD"),
		APP_URL: os.Getenv("APP_URL"),
	}

	return config, err
}
