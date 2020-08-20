package mw

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func failedMetricRegistry(name string, err error, logger *log.Logger) {
	if err != nil {
		logger.WithFields(log.Fields{
			"caller":    "kernel.repository.factory",
			"collector": strings.ToLower(name),
			"kind":      "prometheus",
			"detail":    err.Error(),
		}).Warnf("metric collector '%s' registry failed", strings.ToLower(name))
	}
}
