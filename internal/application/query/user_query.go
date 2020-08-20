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
	ctxI, _ := context.WithCancel(ctx)
	users, nextToken, err = q.repository.Fetch(ctxI, *criteria)
	return
}
