package command

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
)

type UserCommandImp struct {
	repository domain.UserCommandRepository
}

func NewUserCommand(r domain.UserCommandRepository) *UserCommandImp {
	return &UserCommandImp{
		repository: r,
	}
}

func (c *UserCommandImp) Enable(ctx context.Context, id string) error {
	ctxI, _ := context.WithCancel(ctx)
	return c.repository.Restore(ctxI, id)
}

func (c *UserCommandImp) Disable(ctx context.Context, id string) error {
	ctxI, _ := context.WithCancel(ctx)
	return c.repository.Remove(ctxI, id)
}

func (c *UserCommandImp) Remove(ctx context.Context, id string) error {
	ctxI, _ := context.WithCancel(ctx)
	return c.repository.HardRemove(ctxI, id)
}
