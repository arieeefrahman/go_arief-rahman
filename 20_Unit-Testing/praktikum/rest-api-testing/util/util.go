package util

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile("D:/Coding/MSIB/Alterra/go_arief-rahman/20_Unit-Testing/praktikum/rest-api-testing/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config %s", err)
	}

	return viper.GetString(key)
}