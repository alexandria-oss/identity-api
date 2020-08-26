package cmdhandler

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
)

// Contains invokers
type UserHandlerImp struct {
	repository repository.User
}

func NewUserCommandHandler(r repository.User) *UserHandlerImp {
	return &UserHandlerImp{
		repository: r,
	}
}

func (c *UserHandlerImp) Enable(cmd command.Enable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Restore(ctxI, cmd.ID)
}

func (c *UserHandlerImp) Disable(cmd command.Disable) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.Remove(ctxI, cmd.ID)
}

func (c *UserHandlerImp) Remove(cmd command.Remove) error {
	ctxI, _ := context.WithCancel(cmd.Ctx)
	return c.repository.HardRemove(ctxI, cmd.ID)
}
