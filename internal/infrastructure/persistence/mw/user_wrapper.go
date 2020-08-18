package mw

import (
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
)

// HoC-like repository wrapping using the chain of responsibility pattern.
// This wraps required layers to any repository's endpoint such as:
//
// - Observability (Metrics and Tracing)
//
// - Query Caching
//
// - Resiliency strategies (Circuit Breaking and Retry Policy)
func WrapUserRepository(r repository.User, p *redis.Client, k domain.KernelStore, l *log.Logger) repository.User {
	// Chain order: Tracing -> Metric -> Cache (if available) -> Repository
	var repo repository.User
	repo = r
	if p != nil {
		repo = UserRepositoryCache{
			DB:     p,
			Kernel: k,
			Logger: l,
			Mu:     new(sync.RWMutex),
			Next:   repo,
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
