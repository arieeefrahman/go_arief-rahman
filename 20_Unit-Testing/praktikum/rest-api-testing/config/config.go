package config

import "prak-testing/util"

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port string
	DB_Host string
	DB_Name string
}

func InitConfig() *Config {
	defaultConfig := Config{
		DB_Username: util.GetConfig("DB_USERNAME"),
		DB_Password: util.GetConfig("DB_PASSWORD"),
		DB_Port: util.GetConfig("DB_PORT"),
		DB_Host: util.GetConfig("DB_HOST"),
		DB_Name: util.GetConfig("DB_NAME"),
	}

	return &defaultConfig
}

func InitTestConfig() *Config {
	defaultConfig := Config{
		DB_Username: util.GetConfig("DB_USERNAME"),
		DB_Password: util.GetConfig("DB_PASSWORD"),
		DB_Port: util.GetConfig("DB_PORT"),
		DB_Host: util.GetConfig("DB_HOST"),
		DB_Name: util.GetConfig("DB_TEST_NAME"),
	}

	return &defaultConfig
}

