package service

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/common"
	"github.com/alexandria-oss/identity-api/internal/domain"
)

type UserQuery interface {
	Get(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context, token string, size int, filterMap common.FilterMap) ([]*domain.User, common.QueryToken, error)
}
