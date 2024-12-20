package middleware

import "github.com/rs/zerolog/log"

func AuthMiddleware() {
	log.Info().Msg("Doing something in another file")
	log.Error().Err(nil).Msg("This is a simulated error")
}
