package mw

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"time"
)

type UserRepositoryCache struct {
	Repo repository.Cache
	Next repository.User
}

var userTable = "user"

// Write -> Invalidate
// Read -> Cache aside/Write through pattern(s)

func (u UserRepositoryCache) Remove(ctx context.Context, id string) (err error) {
	err = u.Next.Remove(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		_ = u.Repo.Invalidate(ctx, userTable, id)
	}
	return
}

func (u UserRepositoryCache) Restore(ctx context.Context, id string) (err error) {
	err = u.Next.Restore(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		_ = u.Repo.Invalidate(ctx, userTable, id)
	}
	return
}

func (u UserRepositoryCache) HardRemove(ctx context.Context, id string) (err error) {
	err = u.Next.HardRemove(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		_ = u.Repo.Invalidate(ctx, userTable, id)
	}
	return
}

func (u UserRepositoryCache) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	// Cache-aside
	if res, errR := u.Repo.Read(ctx, userTable, key); errR == nil && res != "" {
		user = new(aggregate.UserRoot)
		if errJSON := user.UnmarshalBinary([]byte(res)); errJSON == nil {
			return
		}
	}

	user, err = u.Next.FetchOne(ctx, byUsername, key)
	// Write-through
	_ = u.Repo.Write(ctx, userTable, key, user, time.Minute*30)
	return
}

func (u UserRepositoryCache) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot, err error) {
	// Hash params and queries into single string -> KEY = service:hashed_query - VAL []UserRoot (cache query strategy)
	encodedCriteria := u.hashCriteria(criteria)
	if res, errR := u.Repo.Read(ctx, userTable, encodedCriteria); errR == nil && res != "" {
		users = make([]*aggregate.UserRoot, 0)
		if errJSON := json.Unmarshal([]byte(res), &users); errJSON == nil {
			return
		}
	}

	// If not available, execute query in real db
	users, err = u.Next.Fetch(ctx, criteria)
	// Apply write through strategy
	if usersJSON, errJSON := json.Marshal(users); errJSON == nil {
		_ = u.Repo.Write(ctx, userTable, encodedCriteria, usersJSON, time.Minute*10)
	}
	return
}

/* Helper functions */

func (u UserRepositoryCache) hashCriteria(c domain.Criteria) string {
	h := crypto.SHA1.New()
	criteriaJSON, err := c.MarshalBinary()
	if err != nil {
		return ""
	}
	h.Write(criteriaJSON)

	return hex.EncodeToString(h.Sum(nil))
}
