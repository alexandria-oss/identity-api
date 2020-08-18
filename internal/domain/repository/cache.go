package repository

import (
	"context"
	"time"
)

type Cache interface {
	Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) error
	Read(ctx context.Context, table, key string) (string, error)
	Invalidate(ctx context.Context, table, key string) error
}
