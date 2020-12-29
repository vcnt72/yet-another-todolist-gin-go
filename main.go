package main

import (
	"github.com/yet-another-todo-list-golang/api"
	"github.com/yet-another-todo-list-golang/config"
)

func main() {
	server := api.Server()
	err := server.Run(":" + config.GetEnvConfig("server.port"))
	if err != nil {
		panic(err.Error())
	}
}
