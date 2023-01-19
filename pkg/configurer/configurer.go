package configurer

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// MapConfig It reads in the configuration file and maps it to the interface(input struct).
func MapConfig(conf interface{}) error {

	viper.SetDefault("config.path", "./config")
	err := viper.BindEnv("config.path", "CONFIG_PATH")
	if err != nil {
		log.Printf("warning: %s \n", err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(viper.GetString("config.path"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("warning: %s \n", err)
		return err
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}
