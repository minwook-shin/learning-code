package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("Print")

	log.Info().Str("strKey", "strVal").Msg("Msg")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Info")

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Str("strKey", "strVal").Msg("Msg")

	logger = log.With().Str("strKey", "strVal").Logger()
	logger.Info().Msg("Msg")

	log.Info().Str("strKey", "strVal").
		Dict("dict", zerolog.Dict().
			Str("strKey", "strVal"),
		).Msg("Msg")

	sampled := log.Sample(&zerolog.BasicSampler{N: 10})
	sampled.Info().Msg("Msg")

	log.Logger = log.With().Caller().Logger()
	log.Info().Msg("Msg")

	zerolog.TimestampFieldName = "TimestampFieldName"
	zerolog.LevelFieldName = "LevelFieldName"
	zerolog.MessageFieldName = "MessageFieldName"
	log.Info().Msg("Msg")

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("strKey", "strVal").Msg("Msg")
}
