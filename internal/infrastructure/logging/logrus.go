package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var logger *log.Logger
var loggerSingleton = new(sync.Once)

func NewLogger() *log.Logger {
	if logger == nil {
		loggerSingleton.Do(func() {
			logger = log.New()
			logger.SetFormatter(&log.JSONFormatter{})
			logger.SetOutput(os.Stdout)
		})
	}

	return logger
}
