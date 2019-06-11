package main

import (
	"fmt"
	"weeny/apiserver"
	"weeny/application"
	"weeny/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("== weeny ==")

	c := config.CreateCacheFromConfig()
	h := config.CreateHasherFromConfig()

	s := apiserver.NewServer(
		application.New(
			c,
			h,
		),
	)

	log.Info().Str("cache", fmt.Sprintf("%T", c)).Str("hasher", fmt.Sprintf("%T", h)).Msg("Initializing")

	if err := s.Start(8000); err != nil {
		log.Error().Msgf("Initializing server: %s", err)
	}
}
