package mw

import (
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	log "github.com/sirupsen/logrus"
)

// HoC-like repository wrapping using the chain of responsibility pattern.
// This wraps required layers to each repository's endpoint such as:
//
// - Observability (Logging, Metrics and Tracing)
//
// - Resiliency strategies (Circuit Breaking and Retry Policy)
func WrapCacheRepository(c repository.Cache, l *log.Logger) repository.Cache {
	// Chain order: Tracing -> Metric -> Logging -> Repository
	var repo repository.Cache
	repo = c
	repo = CacheLog{
		Logger: l,
		Next:   repo,
	}
	repo = NewCacheMetric(repo, l)
	repo = CacheTracing{
		Next: repo,
	}

	return repo
}
