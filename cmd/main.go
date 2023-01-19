package main

import (
	"fmt"

	"github.com/Napat/go_loadconfig_sample/internal/config"
	"github.com/Napat/go_loadconfig_sample/pkg/configurer"
	"github.com/pkg/errors"
)

func main() {

	conf := &config.AppConfig{}
	if err := configurer.MapConfig(conf); err != nil {
		panic(errors.Wrap(err, "load app config error"))
	}
	fmt.Printf("conf before secret: %#v\n", conf)

	configurer.LoadSecretAppEnvFile(&conf.Secrets, "./config/secret.env")
	fmt.Printf("conf after secret: %#v\n", conf)
}
