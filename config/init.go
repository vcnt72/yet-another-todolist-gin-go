package config

import "github.com/yet-another-todo-list-golang/db"

func init() {
	initEnv()
	db.GetConnection()
}
