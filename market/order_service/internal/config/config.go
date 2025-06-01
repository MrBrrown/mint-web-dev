package config

import (
	"log"
	"market/common/storage"
	"market/common/yamlconf"
)

type Config struct {
	DB      storage.DbInfo `yaml:"db"`
	BinAddr string         `yaml:"bin_addr"`

	JWTSecret string `yaml:"secret"`
}

func New(path string) *Config {
	cfg := &Config{}
	err := yamlconf.Load(path, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
