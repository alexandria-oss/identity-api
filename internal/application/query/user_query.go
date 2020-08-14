package query

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
)

type UserQueryImp struct {
	repository repository.User
}

func NewUserQuery(r repository.User) *UserQueryImp {
	return &UserQueryImp{
		repository: r,
	}
}

func (q *UserQueryImp) Get(ctx context.Context, username string) (*entity.User, error) {
	ctxI, _ := context.WithCancel(ctx)
	return q.repository.FetchOne(ctxI, true, username)
}

func (q *UserQueryImp) List(ctx context.Context, criteria domain.Criteria) (users []*entity.User,
	nextToken domain.PaginationToken, err error) {
	// Request next token
	nextSize := criteria.Limit + 1

	ctxI, _ := context.WithCancel(ctx)
	users, err = q.repository.Fetch(ctxI, criteria.Token.GetPrimitive(), nextSize.GetPrimitive(), criteria.FilterBy)
	if err != nil {
		return
	}

	if criteria.Limit.GetPrimitive() <= len(users) {
		nextToken = domain.PaginationToken(users[len(users)-1].Sub)
		users = users[:len(users)-1]
	}

	return
}
