package main

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
)

type (
	pgConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		SSLMode  string `default:"disable"`
	}

	serverConfig struct {
		Debug bool
		Addr  string
	}

	configSchema struct {
		Server serverConfig
		PG     pgConfig
	}
)

func loadConfig(filename string) (*configSchema, error) {
	cfg := &configSchema{}

	cf := configor.New(&configor.Config{
		AutoReload: false,
		ENVPrefix:  "",
	})

	//if config.yml exists, load it
	if _, err := os.Stat(filename); err == nil {
		log.Debug().Msgf("loading config file: %s", filename)
		if err := cf.Load(cfg, filename); err != nil {
			return cfg, err
		}
		return cfg, nil
	}

	log.Debug().Msg("loading config from os env")
	if err := cf.Load(cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
