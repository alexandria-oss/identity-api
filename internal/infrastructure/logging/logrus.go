package logging

import (
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"sync"
	"time"
)

var logger *log.Logger
var loggerSingleton = new(sync.Once)

func NewLogger(kernel domain.KernelStore) *log.Logger {
	if logger == nil {
		loggerSingleton.Do(func() {
			filename := "log/alexandria.log"
			level := log.DebugLevel

			if kernel.Environment == domain.Production {
				filename = "/var/log/alexandria.log"
				level = log.InfoLevel
			}

			rotateHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
				Filename:   filename,
				MaxSize:    50,
				MaxBackups: 3,
				MaxAge:     28,
				Level:      level,
				Formatter: &log.JSONFormatter{
					TimestampFormat: time.RFC822,
				},
			})

			logger = log.New()
			logger.SetLevel(level)
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
