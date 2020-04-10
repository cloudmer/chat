package tool

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var once sync.Once
var Logger zerolog.Logger

func init()  {
	once.Do(func() {
		Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		// print code file row number
		//Logger = Logger.With().Caller().Logger()
		Logger = Logger.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
			TimeFormat: TimeFormat,
		})
	})
}

