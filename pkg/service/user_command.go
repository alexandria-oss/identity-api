package service

import "context"

type UserCommand interface {
	Enable(ctx context.Context, id string) error
	Disable(ctx context.Context, id string) error
	Remove(ctx context.Context, id string) error
}
