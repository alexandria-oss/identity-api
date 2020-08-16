package mw

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type UserRepositoryCache struct {
	DB     *redis.Client
	Kernel domain.KernelStore
	Logger *log.Logger
	Mu     *sync.RWMutex
	Next   repository.User
}

// Write -> Invalidate
// Read -> Cache aside/Write through pattern(s)

func (u UserRepositoryCache) Remove(ctx context.Context, id string) (err error) {
	u.Mu.Lock()
	defer u.Mu.Unlock()

	err = u.Next.Remove(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		key := fmt.Sprintf("%s:%s", u.Kernel.Service, id)
		if conn := u.DB.Conn(ctx); conn != nil {
			defer func() {
				_ = conn.Close()
			}()
			errDel := conn.Del(ctx, key).Err()
			if errDel != nil {
				u.Logger.WithFields(log.Fields{
					"key":    key,
					"detail": errDel.Error(),
				}).Error("failed to invalidate user cache")
			}
		}
	}
	return
}

func (u UserRepositoryCache) Restore(ctx context.Context, id string) (err error) {
	u.Mu.Lock()
	defer u.Mu.Unlock()

	err = u.Next.Restore(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		key := fmt.Sprintf("%s:%s", u.Kernel.Service, id)
		if conn := u.DB.Conn(ctx); conn != nil {
			defer func() {
				_ = conn.Close()
			}()
			errDel := conn.Del(ctx, key).Err()
			if errDel != nil {
				u.Logger.WithFields(log.Fields{
					"key":    key,
					"detail": errDel.Error(),
				}).Error("failed to invalidate user cache")
			}
		}
	}
	return
}

func (u UserRepositoryCache) HardRemove(ctx context.Context, id string) (err error) {
	u.Mu.Lock()
	defer u.Mu.Unlock()

	err = u.Next.HardRemove(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		key := fmt.Sprintf("%s:%s", u.Kernel.Service, id)
		if conn := u.DB.Conn(ctx); conn != nil {
			defer func() {
				_ = conn.Close()
			}()
			errDel := conn.Del(ctx, key).Err()
			if errDel != nil {
				u.Logger.WithFields(log.Fields{
					"key":    key,
					"detail": errDel.Error(),
				}).Error("failed to invalidate user cache")
			}
		}
	}
	return
}

func (u UserRepositoryCache) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	u.Mu.RLock()
	defer u.Mu.RUnlock()

	// Cache-aside
	keyMem := fmt.Sprintf("%s:%s", u.Kernel.Service, key)
	if conn := u.DB.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		if res, errMem := conn.Get(ctx, keyMem).Result(); errMem == nil && len(res) > 1 {
			user = new(aggregate.UserRoot)
			if errJSON := user.UnmarshalBinary([]byte(res)); errJSON == nil {
				return
			}
		}
	}

	user, err = u.Next.FetchOne(ctx, byUsername, key)
	// Write-through
	if conn := u.DB.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		if err == nil && user != nil {
			if errSet := conn.Set(ctx, keyMem, user, time.Minute*30).Err(); errSet != nil {
				u.Logger.WithFields(log.Fields{
					"key":    key,
					"detail": errSet.Error(),
				}).Error("failed to write cache")
			}
		}
	}
	return
}

func (u UserRepositoryCache) Fetch(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*aggregate.UserRoot, error) {
	u.Mu.RLock()
	defer u.Mu.RUnlock()

	return u.Next.Fetch(ctx, token, size, filterMap)
}
