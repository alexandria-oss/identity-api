package repository

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
)

type User interface {
	Remove(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	HardRemove(ctx context.Context, id string) error
	FetchOne(ctx context.Context, byUsername bool, key string) (*entity.User, error)
	Fetch(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*entity.User, error)
}
