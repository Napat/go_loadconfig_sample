package main

import (
	"fmt"

	"github.com/Napat/go_loadconfig_sample/internal/config"
	"github.com/Napat/go_loadconfig_sample/pkg/configurer"
	"github.com/pkg/errors"
)

func main() {

	conf := new(config.AppConfig)
	if err := configurer.LoadConfig(conf, "./config", "config", "yaml", "APPENV"); err != nil {
		panic(errors.Wrap(err, "load app config error"))
	}
	fmt.Printf("conf before secret: %#v\n", conf)

	configurer.LoadDotEnv(&conf.Secrets, "./config/secret.env", "SECRET", "APPENV")
	fmt.Printf("conf after secret: %#v\n", conf)
}
