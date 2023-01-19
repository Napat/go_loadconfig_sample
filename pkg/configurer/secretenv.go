package configurer

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// LoadSecretAppEnvFile It loads an environment file based on the value of APPENV in the current environment.
func LoadSecretAppEnvFile(secrets interface{}, secretPath string) {

	loadSecretAppEnvFile(secretPath)

	// env to struct
	err := envconfig.Process("", secrets)
	if err != nil {
		log.Fatal("Error unmarshalling env vars")
	}
}

func loadSecretAppEnvFile(secretPath string) {
	currentEnvironment, ok := os.LookupEnv("APPENV")
	if ok {
		fmt.Printf("loadSecretAppEnvFile: %v\n", currentEnvironment)
		if err := godotenv.Load(secretPath + "." + currentEnvironment); err != nil {
			panic(err)
		}
	}
}
