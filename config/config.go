package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type config struct {
	appName       string
	port          int
	migrationPath string
	db            dbConfig
}

var appConfig config

func Load() {
	viper.SetDefault("APP_NAME", "boilerplate")
	viper.SetDefault("APP_PORT", "8000")

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	appConfig = config{
		appName:       readEnvString("APP_NAME"),
		port:          readEnvInt("APP_PORT"),
		migrationPath: readEnvString("MIGRATION_PATH"),
		db:            getDatabaseConfig(),
	}
}

func AppPort() int {
	return appConfig.port
}

func AppName() string {
	return appConfig.appName
}

func Migrationpath() string {
	return appConfig.migrationPath
}

func readEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid number", key))
	}
	return v
}

func readEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := errors.New(fmt.Sprintf("Key %s is not set", key))
		panic(err)
	}
}
