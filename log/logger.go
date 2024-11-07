/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var FlagDebug bool

func InitializeLogger() {
	_noColor := !viper.GetBool("logging.colored_output")

	if os.Getenv("INVOCATION_ID") == "" {
		// detect systemd
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: _noColor, TimeFormat: time.RFC3339})
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, PartsExclude: []string{"time"}})
	}

	if FlagDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
