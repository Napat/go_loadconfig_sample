package configurer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// LoadDotEnv It loads an environment file to struct based on the value of osenv.
// ie. LoadDotEnv(&conf.Secrets, "./config/secret.env", "SECRET", "APPENV")
func LoadDotEnv(secrets interface{}, secretPathName string, secretPrefix string, osenv string) error {
	if osenv == "" {
		osenv = "APPENV"
	}

	currentEnvironment, ok := os.LookupEnv(osenv)
	if ok {
		if currentEnvironment != "local" {
			secretPathName = secretPathName + "." + currentEnvironment
		}
		log.Printf("LoadDotEnv env: %s, secretPathName: %s\n", currentEnvironment, secretPathName)
		if err := godotenv.Load(secretPathName); err != nil {
			panic(err)
		}
	}

	// env to struct
	err := envconfig.Process(secretPrefix, secrets)
	if err != nil {
		panic("Error unmarshalling env vars")
	}
	return nil
}
