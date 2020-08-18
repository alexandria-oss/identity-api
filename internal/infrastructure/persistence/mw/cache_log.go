package mw

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type CacheLog struct {
	Logger *log.Logger
	Next   repository.Cache
}

func (c CacheLog) Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) (err error) {
	defer func() {
		if err != nil {
			c.Logger.WithFields(log.Fields{
				"caller": "kernel.repository.cache",
				"table":  table,
				"key":    key,
				"detail": err.Error(),
			}).Error("failed to write cache")
		} else {
			c.Logger.WithFields(log.Fields{
				"caller": "kernel.repository.cache",
				"table":  table,
				"key":    key,
			}).Info("cached value")
		}
	}()
	err = c.Next.Write(ctx, table, key, value, duration)
	return
}

func (c CacheLog) Read(ctx context.Context, table, key string) (string, error) {
	return c.Next.Read(ctx, table, key)
}

func (c CacheLog) Invalidate(ctx context.Context, table, key string) (err error) {
	defer func() {
		if err != nil {
			c.Logger.WithFields(log.Fields{
				"caller": "kernel.repository.cache",
				"table":  table,
				"key":    key,
				"detail": err.Error(),
			}).Error("failed to invalidate cache")
		}
	}()

	err = c.Next.Invalidate(ctx, table, key)
	return
}
