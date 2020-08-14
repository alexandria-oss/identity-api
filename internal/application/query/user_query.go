package query

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/user"
)

type UserQueryImp struct {
	repository user.UserQueryRepository
}

func NewUserQuery(r user.UserQueryRepository) *UserQueryImp {
	return &UserQueryImp{
		repository: r,
	}
}

func (q *UserQueryImp) Get(ctx context.Context, username string) (*user.User, error) {
	ctxI, _ := context.WithCancel(ctx)
	return q.repository.FetchOne(ctxI, true, username)
}

func (q *UserQueryImp) List(ctx context.Context, token string, size int, filterMap domain.FilterMap) (users []*user.User,
	nextToken domain.QueryToken, err error) {
	// Request next token
	nextSize := size + 1

	ctxI, _ := context.WithCancel(ctx)
	users, err = q.repository.Fetch(ctxI, token, nextSize, filterMap)
	if err != nil {
		return
	}

	if size <= len(users) {
		nextToken = domain.QueryToken(users[len(users)-1].Sub)
		users = users[:len(users)-1]
	}

	return
}
