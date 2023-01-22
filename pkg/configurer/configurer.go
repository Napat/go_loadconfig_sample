package configurer

import (
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// LoadConfig It reads in the configuration file and maps it to the input struct interface.
func LoadConfig(conf interface{}, configPath, configName, configType, osenv string) error {
	currentEnvironment, ok := os.LookupEnv(osenv)
	if ok && (currentEnvironment != "local") {
		configName = configName + "." + currentEnvironment
	}

	viper.SetDefault("config.path", configPath)
	if err := viper.BindEnv("config.path", "CONFIG_PATH"); err != nil {
		log.Printf("warning: %s \n", err)
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(viper.GetString("config.path"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config error")
	}

	if err := viper.Unmarshal(conf); err != nil {
		return errors.Wrap(err, "unmarshal config to struct error")
	}
	return nil
}
