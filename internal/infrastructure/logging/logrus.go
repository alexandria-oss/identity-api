package logging

import (
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"sync"
	"time"
)

var logger *log.Logger
var loggerSingleton = new(sync.Once)

func NewLogger() *log.Logger {
	if logger == nil {
		loggerSingleton.Do(func() {
			rotateHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
				Filename:   "/var/log/alexandria.log",
				MaxSize:    50,
				MaxBackups: 3,
				MaxAge:     28,
				Level:      log.InfoLevel,
				Formatter: &log.JSONFormatter{
					TimestampFormat: time.RFC822,
				},
			})

			logger = log.New()
			logger.SetLevel(log.InfoLevel)
			logger.SetOutput(colorable.NewColorableStdout())
			logger.SetFormatter(&log.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: time.RFC822,
			})

			if err != nil {
				logger.Errorf("failed to initialize file rotate hook: %v", err)
				return
			}

			logger.AddHook(rotateHook)
		})
	}

	return logger
}
