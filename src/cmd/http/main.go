package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"weeny/apiserver"
	"weeny/application"
	"weeny/config"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("== weeny ==")

	s := apiserver.NewServer(
		application.New(
			config.CreateCacheFromConfig(),
			config.CreateHasherFromConfig(),
		),
	)

	if err := s.Start(8000); err != nil {
		log.Error().Msgf("Initializing server: %s", err)
	}
}
