package internal

import "github.com/rs/zerolog/log"

func GetMessage(s string) string {
	defer log.Info().Msg("Returned Message")
	return "Hello " + s
}