package domain

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/common"
)

type UserQueryRepository interface {
	FetchOne(ctx context.Context, byUsername bool, key string) (*User, error)
	Fetch(ctx context.Context, token string, size int, filterMap common.FilterMap) ([]*User, error)
}
