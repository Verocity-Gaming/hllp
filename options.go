package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Options represents the user configurations created on startup.
var (
	// Logging options.
	verbosity string
	color     bool

	certificate string
	key         string

	port string

	addr     string
	password string
)

// parse will return an Options from parsed command line flags
func parse() error {
	flag.StringVar(&verbosity, "v", zerolog.InfoLevel.String(), "logging verbosity level [panic, error, warn, info, debug]")
	flag.BoolVar(&color, "color", true, "logging with color")

	flag.StringVar(&certificate, "cert", "", "/etc/ssl/certs/server.crt")
	flag.StringVar(&key, "key", "", "/etc/ssl/private/server.key")

	flag.StringVar(&port, "port", "", "port to listen on")
	flag.StringVar(&addr, "addr", "", "RCON address")
	flag.StringVar(&password, "pass", "", "RCON password")

	flag.Parse()

	l, err := zerolog.ParseLevel(verbosity)
	if err != nil {
		return err
	}

	log.Logger = zerolog.New(os.Stderr).
		Level(l).
		With().
		Timestamp().
		Logger()

	log.Info().
		Str("verbosity", verbosity).
		Bool("color", color).
		Str("cert", certificate).
		Str("key", key).
		Str("port", port).
		Str("rcon", addr).
		Msg("starting HLLP server")

	return nil
}
