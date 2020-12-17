package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//InitEnv env initialize config wrapper
func InitEnv() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error on ENV")
		os.Exit(1)
	}

	log.Println("Environment is being loaded successfully")
}
