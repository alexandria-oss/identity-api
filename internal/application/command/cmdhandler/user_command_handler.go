package cmdhandler

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
)

// Contains invokers
type UserImp struct {
	repository repository.User
}

func NewUserCommandHandler(r repository.User) *UserImp {
	return &UserImp{
		repository: r,
	}
}

func (c *UserImp) Enable(cmd command.Enable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Restore(ctxI, cmd.ID)
}

func (c *UserImp) Disable(cmd command.Disable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Remove(ctxI, cmd.ID)
}

func (c *UserImp) Remove(cmd command.Remove) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.HardRemove(ctxI, cmd.ID)
}
