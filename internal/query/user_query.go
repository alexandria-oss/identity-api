package query

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/common"
	"github.com/alexandria-oss/identity-api/internal/domain"
)

type UserQueryImp struct {
	repository domain.UserQueryRepository
}

func NewUserQueryImp(r domain.UserQueryRepository) *UserQueryImp {
	return &UserQueryImp{
		repository: r,
	}
}

func (q *UserQueryImp) Get(ctx context.Context, username string) (*domain.User, error) {
	ctxI, _ := context.WithCancel(ctx)
	return q.repository.FetchOne(ctxI, true, username)
}

func (q *UserQueryImp) List(ctx context.Context, token string, size int, filterMap common.FilterMap) (users []*domain.User,
	nextToken common.QueryToken, err error) {
	// Request next token
	nextSize := size + 1

	ctxI, _ := context.WithCancel(ctx)
	users, err = q.repository.Fetch(ctxI, token, nextSize, filterMap)
	if err != nil {
		return
	}

	if size <= len(users) {
		nextToken = common.QueryToken(users[len(users)-1].Sub)
		users = users[:len(users)-1]
	}

	return
}
