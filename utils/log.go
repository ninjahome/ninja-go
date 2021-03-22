package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"sync"
	"time"
)

var _logInstance *zerolog.Logger
var logOnce sync.Once

func LogInst() *zerolog.Logger {
	logOnce.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339Nano
		writer := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
		logger := zerolog.New(zerolog.ConsoleWriter{Out: writer}).
			Level(config.LogLevel).
			With().
			Caller().
			Timestamp().
			Logger()
		_logInstance = &logger
	})

	return _logInstance
}
