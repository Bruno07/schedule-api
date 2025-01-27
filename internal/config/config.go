package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	
}

func GetPort() string {
	return fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}