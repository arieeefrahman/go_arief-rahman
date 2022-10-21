package util

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	
	if err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	resp := viper.GetString(key)

	return resp
}