package logger

import (
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init sets up a global logger instance
func InitLogger() {
	logLevel := strings.ToUpper("DEBUG") // hardcoded TODO: add env var
	switch logLevel {
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Info().Msgf("setting global log level to %s", logLevel)
}

func LogInfo(msg string) {
	if msg != "" {
		log.Info().Msg(msg)
	}
}
