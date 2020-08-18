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
	// Chain order: Tracing -> Metric -> Cache (if available) -> Repository
	var repo repository.User
	repo = r
	if c != nil {
		repo = UserRepositoryCache{
			Repo: c,
			Next: repo,
		}
	}
	repo = UserRepositoryMetric{
		Next: repo,
	}
	repo = UserRepositoryTracing{
		Next: repo,
	}

	return repo
}
