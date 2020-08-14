package service

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/user"
)

type UserQuery interface {
	Get(ctx context.Context, id string) (*user.User, error)
	List(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*user.User, domain.QueryToken, error)
}
