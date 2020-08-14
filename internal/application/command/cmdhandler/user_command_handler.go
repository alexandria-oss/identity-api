package cmdhandler

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
)

// Contains invokers
type UserCommandHandlerImp struct {
	repository repository.User
}

func NewUserCommandHandler(r repository.User) *UserCommandHandlerImp {
	return &UserCommandHandlerImp{
		repository: r,
	}
}

func (c *UserCommandHandlerImp) Enable(cmd command.Enable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Restore(ctxI, cmd.ID)
}

func (c *UserCommandHandlerImp) Disable(cmd command.Disable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Remove(ctxI, cmd.ID)
}

func (c *UserCommandHandlerImp) Remove(cmd command.Remove) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.HardRemove(ctxI, cmd.ID)
}
