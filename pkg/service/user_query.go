package service

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
)

type UserQuery interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*entity.User, domain.QueryToken, error)
}
