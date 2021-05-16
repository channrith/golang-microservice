package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const DB_URI = "federation_user:rN1D7F3G@tcp(127.0.0.1:3306)/tc_online_db"

func DotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}
