package query

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
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

func (q *UserQueryImp) Get(ctx context.Context, username string) (*aggregate.UserRoot, error) {
	ctxI, _ := context.WithCancel(ctx)
	return q.repository.FetchOne(ctxI, true, username)
}

func (q *UserQueryImp) GetByID(ctx context.Context, id string) (*aggregate.UserRoot, error) {
	ctxI, _ := context.WithCancel(ctx)
	return q.repository.FetchOne(ctxI, false, id)
}

func (q *UserQueryImp) List(ctx context.Context, criteria *domain.Criteria) (users []*aggregate.UserRoot,
	nextToken domain.PaginationToken, err error) {
	// Request next token
	criteria.Limit = criteria.Limit + 1

	ctxI, _ := context.WithCancel(ctx)
	users, err = q.repository.Fetch(ctxI, *criteria)
	if err != nil {
		return
	}

	if criteria.Limit.GetPrimitive() <= len(users) {
		nextToken = domain.PaginationToken(users[len(users)-1].Root.Sub)
		users = users[:len(users)-1]
	}

	return
}
