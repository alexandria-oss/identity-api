package user

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
)

type UserQueryRepository interface {
	FetchOne(ctx context.Context, byUsername bool, key string) (*User, error)
	Fetch(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*User, error)
}
