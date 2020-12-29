package config

import "log"

func Init() {
	log.Println("Initiating configuration")
	initEnv()
}
