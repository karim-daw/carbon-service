package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get env file
func HasEnvFile() error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return err
	}
	log.Print("Found .env file")
	return nil
}

// returns the env variable given string key
func LoadEnvVar(key string) string {
	envVariable, exists := os.LookupEnv(key)
	if !exists {
		fmt.Printf("Could not find %s \n", key)
	}
	return envVariable
}
