package mw

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
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
		key := u.generateComposedKey(u.Kernel.Service, id)
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
		key := u.generateComposedKey(u.Kernel.Service, id)
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
		key := u.generateComposedKey(u.Kernel.Service, id)
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
	keyMem := u.generateComposedKey(u.Kernel.Service, key)
	if conn := u.DB.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		if res, errMem := conn.Get(ctx, keyMem).Result(); errMem == nil && res != "" {
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

func (u UserRepositoryCache) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot, err error) {
	u.Mu.RLock()
	defer u.Mu.RUnlock()

	// Hash params and queries into single string -> KEY = service:hashed_query - VAL []UserRoot (cache query strategy)
	encodedCriteria := u.hashCriteria(criteria)
	key := u.generateComposedKey(u.Kernel.Service, encodedCriteria)
	if conn := u.DB.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()

		if res, errGet := conn.Get(ctx, key).Result(); errGet == nil && res != "" {
			users = make([]*aggregate.UserRoot, 0)
			if errJSON := json.Unmarshal([]byte(res), &users); errJSON == nil {
				return
			}
		}
	}

	// If not available, execute query in real db
	users, err = u.Next.Fetch(ctx, criteria)
	// Apply write through strategy
	if err == nil {
		if conn := u.DB.Conn(ctx); conn != nil {
			defer func() {
				_ = conn.Close()
			}()

			if usersJSON, errJSON := json.Marshal(users); errJSON == nil {
				if errSet := conn.Set(ctx, key, usersJSON, time.Minute*15).Err(); errSet != nil {
					u.Logger.WithFields(log.Fields{
						"key":    key,
						"detail": errSet.Error(),
					}).Error("failed to write cache")
				}
			}
		}
	}
	return
}

/* Helper functions */

func (u UserRepositoryCache) generateComposedKey(primary, secondary string) string {
	return fmt.Sprintf("%s:%s", primary, secondary)
}

func (u UserRepositoryCache) hashCriteria(c domain.Criteria) string {
	h := crypto.SHA1.New()
	criteriaJSON, err := c.MarshalBinary()
	if err != nil {
		return ""
	}
	h.Write(criteriaJSON)

	return hex.EncodeToString(h.Sum(nil))
}
