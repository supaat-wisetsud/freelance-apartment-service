package utility

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	godotenv.Load()

}

func GetSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
