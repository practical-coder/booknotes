package zlg

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger
var once sync.Once

func init() {
	once.Do(func() {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
		Logger = zerolog.New(multi).With().Timestamp().Logger()
	})
}
