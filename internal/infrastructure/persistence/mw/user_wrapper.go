package mw

import (
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	log "github.com/sirupsen/logrus"
)

// HoC-like repository wrapping using the chain of responsibility pattern.
// This wraps required layers to any repository's endpoint such as:
//
// - Observability (Metrics and Tracing)
//
// - Query Caching
//
// - Resiliency strategies (Circuit Breaking and Retry Policy)
func WrapUserRepository(r repository.User, c repository.Cache, l *log.Logger) repository.User {
	// Chain order: Cache (if available) -> Tracing -> Metric -> Repository
	var repo repository.User
	repo = r
	repo = NewUserRepositoryMetric(repo, l)
	repo = UserRepositoryTracing{
		Next: repo,
	}

	if c != nil {
		// Note: Cache repo contains owned metrics and tracing, and should be kept as a priority/top level
		// to avoid duplicate metrics or tracing spans
		repo = UserRepositoryCache{
			Cache: c,
			Next:  repo,
		}
	}

	return repo
}
