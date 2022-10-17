package config

import "weekly-task-2/util"

type Config struct {
	DB_Username string
	DB_Password string
	DB_Name     string
	DB_Host     string
	DB_Port     string
}

func InitConfig() *Config {
	defaultConfig := Config{
		DB_Username: util.GetConfig("DB_USERNAME"),
		DB_Password: util.GetConfig("DB_PASSWORD"),
		DB_Name: util.GetConfig("DB_NAME"),
		DB_Host: util.GetConfig("DB_HOST"),
		DB_Port: util.GetConfig("DB_PORT"),
	}

	return &defaultConfig
}