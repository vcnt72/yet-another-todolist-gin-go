package config

import "fmt"

// databaseConfig database minimal configuration
type databaseConfig struct {
	DBPort     string
	DBHost     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

//getDatabaseConfig generate database config and return the config
func getDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		DBPort:     GetEnvConfig("db.port"),
		DBDatabase: GetEnvConfig("db.database"),
		DBHost:     GetEnvConfig("db.host"),
		DBUsername: GetEnvConfig("db.username"),
		DBPassword: GetEnvConfig("db.password"),
	}
}

//getDsn generate dsn and return the dsn
func GetDsn() string {
	config := getDatabaseConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBDatabase,
		config.DBPort,
	)

	return dsn
}
