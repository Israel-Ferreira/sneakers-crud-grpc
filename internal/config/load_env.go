package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoDbHost     string
	MongoDbPort     string
	MongoDbUsername string
	MongoDbPassword string
)

func LoadEnv() {
	godotenv.Load()

	MongoDbHost = os.Getenv("MONGO_DB_HOST")
	MongoDbPort = os.Getenv("MONGO_DB_PORT")

	MongoDbUsername = os.Getenv("MONGO_DB_USERNAME")
	MongoDbPassword = os.Getenv("MONGO_DB_PASSWORD")
}
