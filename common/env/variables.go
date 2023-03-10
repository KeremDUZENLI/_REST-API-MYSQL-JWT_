package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGO_URL             string
	MONGO_COLLECTION_NAME string

	SECRET_KEY string
)

func Load() {
	godotenv.Load(".env")

	MONGO_URL = os.Getenv("MONGO_URL")
	MONGO_COLLECTION_NAME = os.Getenv("MONGO_COLLECTION_NAME")

	SECRET_KEY = os.Getenv("SECRET_KEY")
}
