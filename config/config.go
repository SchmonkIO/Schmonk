package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fyreek/Schmonk/models"
)

var Config models.Config

func ReadConfig(configfile string) error {
	_, err := os.Open(configfile)
	if err != nil {
		return err
	}
	var config models.Config
	_, err = toml.DecodeFile(configfile, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Config = config
	return err
}
