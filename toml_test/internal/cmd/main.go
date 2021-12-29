package main

import (
	"example/toml_test/internal/config"
	"flag"
	"github.com/rs/zerolog/log"
)

func main() {
	flag.Parse()
	err := config.Init()
	if err != nil {
		log.Err(err).Msg("config.Init()")
	}
	log.Info().Interface("config", config.Conf).Msg("")
}
