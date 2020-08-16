package mw

import (
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
)

// HoC-like repository wrapping with chain of responsibility pattern
func WrapUserRepository(r repository.User, p *redis.Client, k domain.KernelStore, l *log.Logger) repository.User {
	var wrappedRepo repository.User
	wrappedRepo = r
	if p != nil {
		wrappedRepo = UserRepositoryCache{
			DB:     p,
			Kernel: k,
			Logger: l,
			Mu:     new(sync.RWMutex),
			Next:   r,
		}
	}

	return wrappedRepo
}
