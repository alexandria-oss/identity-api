package user

import "context"

type UserCommandRepository interface {
	Remove(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	HardRemove(ctx context.Context, id string) error
}
