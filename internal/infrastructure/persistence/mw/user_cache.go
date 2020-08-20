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
	Cache repository.Cache
	Next  repository.User
}

var userTable = "user"

// Write -> Invalidate
// Read -> Cache aside/Write through pattern(s)

func (u UserRepositoryCache) Remove(ctx context.Context, id string) (err error) {
	defer func() {
		if err == nil {
			_ = u.Cache.Invalidate(ctx, userTable, id)
		}
	}()

	err = u.Next.Remove(ctx, id)
	return
}

func (u UserRepositoryCache) Restore(ctx context.Context, id string) (err error) {
	err = u.Next.Restore(ctx, id)
	if err == nil {
		// Removal done, Invalidate cache
		_ = u.Cache.Invalidate(ctx, userTable, id)
	}
	return
}

func (u UserRepositoryCache) HardRemove(ctx context.Context, id string) (err error) {
	defer func() {
		if err == nil {
			_ = u.Cache.Invalidate(ctx, userTable, id)
		}
	}()

	err = u.Next.HardRemove(ctx, id)
	return
}

func (u UserRepositoryCache) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	// Cache-aside
	if res, errR := u.Cache.Read(ctx, userTable, key); errR == nil && res != "" {
		user = new(aggregate.UserRoot)
		if errJSON := user.UnmarshalBinary([]byte(res)); errJSON == nil {
			return
		}
	}

	defer func() {
		if err == nil && user != nil {
			// Write-through
			_ = u.Cache.Write(ctx, userTable, key, user, time.Minute*30)
		}
	}()
	user, err = u.Next.FetchOne(ctx, byUsername, key)
	return
}

func (u UserRepositoryCache) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot,
	nextToken domain.PaginationToken, err error) {
	body := struct {
		Users     []*aggregate.UserRoot  `json:"users"`
		NextToken domain.PaginationToken `json:"next_token"`
	}{
		Users:     make([]*aggregate.UserRoot, 0),
		NextToken: "",
	}
	// Hash params and queries into single string -> KEY = service:hashed_query - VAL []UserRoot (cache query strategy)
	encodedCriteria := u.hashCriteria(criteria)
	if res, errR := u.Cache.Read(ctx, userTable, encodedCriteria); errR == nil && res != "" {
		if errJSON := json.Unmarshal([]byte(res), &body); errJSON == nil {
			users = body.Users
			nextToken = body.NextToken
			return
		}
	}
	defer func() {
		// Write-through
		if err == nil && len(users) > 0 {
			body.Users = users
			body.NextToken = nextToken
			if usersJSON, errJSON := json.Marshal(body); errJSON == nil {
				_ = u.Cache.Write(ctx, userTable, encodedCriteria, usersJSON, time.Minute*10)
			}
		}
	}()

	// If not available, execute query in real db
	users, nextToken, err = u.Next.Fetch(ctx, criteria)
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
