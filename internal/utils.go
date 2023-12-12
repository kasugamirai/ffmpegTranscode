package internal

import (
	"freefrom.space/videoTransform/conf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strings"
	"time"
)

func ConfigZeroLog() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.TimestampFieldName = strings.ToLower("timestamp")
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// check if the file exists
	_, err := os.Stat(conf.GetConf().Zerolog.LogFileName)
	if os.IsNotExist(err) {
		// if not, create the file
		file, err := os.Create(conf.GetConf().Zerolog.LogFileName)
		if err != nil {
			panic(err)
		}
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339Nano})
	} else {
		// file exists, set output target to stderr
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
	}
}
