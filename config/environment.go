package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cast"
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

// GetEnvConfig get key from out of environment package
func GetEnvConfig(key string) string {
	return cast.ToString(viper.Get(key))
}
