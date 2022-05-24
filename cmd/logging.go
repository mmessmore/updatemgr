/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	// Path to log file or "-" for STDERR
	Path string

	// Do pretty logging
	// Also requires Path to be "-"
	Pretty bool

	// Minimum log level to log
	Level string

	// Max number of files to keep
	MaxBackups int

	// size cap in MB
	MaxSize int

	// days to keep
	MaxAge int
}

func ConfigureLogger() {
	lc := LogConfig{
		Path:       viper.GetString("log-file"),
		Pretty:     viper.GetBool("log-fancy"),
		Level:      viper.GetString("log-level"),
		MaxBackups: viper.GetInt("log-backups"),
		MaxSize:    viper.GetInt("log-max-size"),
		MaxAge:     viper.GetInt("log-max-age"),
	}
	lc.Config()
}

func (c *LogConfig) Config() {

	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		// If we're pretty, force STDERR logging
		c.Path = "-"
	}

	// Do this before setting up file logging so the error hits
	// the console
	c.Level = strings.ToUpper(c.Level)
	switch c.Level {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Error().
			Str("value", c.Level).
			Msg("Unrecognized log level, defaulting to INFO")
	}

	// If we're not using STDERR do log rotation
	if c.Path != "-" {
		log.Logger = log.Output(c.newRollingFile())
	}
}

func (c *LogConfig) newRollingFile() io.Writer {
	return &lumberjack.Logger{
		Filename:   c.Path,
		MaxBackups: c.MaxBackups,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
	}
}
