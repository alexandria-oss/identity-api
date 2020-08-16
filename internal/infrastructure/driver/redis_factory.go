package driver

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
)

var redisPool *redis.Client
var redisSingleton = new(sync.Once)

func NewRedisClientPool(ctx context.Context, k domain.KernelStore, logger *log.Logger) (*redis.Client, func()) {
	if redisPool == nil {
		redisSingleton.Do(func() {
			redisPool = redis.NewClient(&redis.Options{
				Network:            k.Config.Cache.Network,
				Addr:               k.Config.Cache.Address[0],
				Dialer:             nil,
				OnConnect:          nil,
				Username:           k.Config.Cache.Username,
				Password:           k.Config.Cache.Password,
				DB:                 k.Config.Cache.Database,
				MaxRetries:         0,
				MinRetryBackoff:    0,
				MaxRetryBackoff:    0,
				DialTimeout:        0,
				ReadTimeout:        0,
				WriteTimeout:       0,
				PoolSize:           0,
				MinIdleConns:       0,
				MaxConnAge:         0,
				PoolTimeout:        0,
				IdleTimeout:        0,
				IdleCheckFrequency: 0,
				TLSConfig:          nil,
				Limiter:            nil,
			})

			if errPing := redisPool.Ping(ctx).Err(); errPing != nil {
				logger.WithFields(log.Fields{
					"address": k.Config.Cache.Address,
				}).Error("failed to connect to redis database")
			}

			logger.WithFields(log.Fields{
				"address": k.Config.Cache.Address,
			}).Info("connected to redis database")
		})
	}

	cleanup := func() {
		if redisPool != nil {
			err := redisPool.Close()
			if err != nil {
				logger.WithField("detail", err.Error()).Error("failed to close redis client connection")
			}
		}
	}

	return redisPool, cleanup
}
