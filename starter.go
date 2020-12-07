package artemis

import (
	"github.com/joho/godotenv"
	"log"
)

func start(env string) {
	envFile := "dev.env"
	if env == "prod" {
		envFile = "prod.env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error Load .env file")
	}
}
