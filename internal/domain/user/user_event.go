package user

import "context"

const (
	Enabled  = "user.enabled"
	Disabled = "user.disabled"
	Removed  = "user.removed"
)

type UserEventBus interface {
	Enabled(ctx context.Context, id string) error
	Disabled(ctx context.Context, id string) error
	Removed(ctx context.Context, id string) error
}
