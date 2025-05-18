package config

import (
	"log"
	"market/common/yamlconf"
)

type Config struct {
}

func New(path string) *Config {
	cfg := &Config{}
	err := yamlconf.Load(path, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
