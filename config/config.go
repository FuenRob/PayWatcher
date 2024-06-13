package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigDataBase struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func GetConfigDataBase() ConfigDataBase {
	/**
	 * TODO: Ver como hacer para que no se tenga que cargar
	 * el archivo .env en cada archivo
	 */
	var config ConfigDataBase
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.User = os.Getenv("DB_USER")
	config.Pass = os.Getenv("DB_PASS")
	config.Name = os.Getenv("DB_NAME")

	return config
}
