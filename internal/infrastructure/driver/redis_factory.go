package driver

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var redisPool *redis.Client
var redisSingleton = new(sync.Once)

func NewRedisClientPool(ctx context.Context, k domain.KernelStore, logger *log.Logger) (*redis.Client, func()) {
	if redisPool == nil {
		// Singleton
		redisSingleton.Do(func() {
			output := make(chan bool, 1)
			// Circuit Breaker
			// Note: CB Timeout policy must be in ms and should be equal to retry policy's total time
			// Retry Policy (x) = (3)(10s)
			// CB Timeout policy f(x) = (x)(1000)   // *result units in ms
			hystrix.DefaultTimeout = int(time.Millisecond * 30000)
			errors := hystrix.GoC(ctx, "redis_pool", func(ctxCB context.Context) error {
				// Retry Policy
				// Actual policy: Try 3 times every 10 sec
				r := retrier.New(retrier.ConstantBackoff(3, time.Second*10), nil)
				err := r.Run(func() error {
					logger.WithFields(log.Fields{
						"caller":  "kernel.data.redis.factory",
						"address": k.Config.Cache.Address,
					}).Info("trying to connect to redis")
					return getRedisPool(ctxCB, k)
				})
				if err != nil {
					return err
				}

				output <- true
				return nil
			}, func(_ context.Context, err error) error {
				return err
			})

			// Wait for Circuit Breaker goroutines to end
			select {
			case <-output:
				logger.WithFields(log.Fields{
					"caller":  "kernel.data.redis.factory",
					"address": k.Config.Cache.Address,
				}).Info("connected to redis database")
				return
			case err := <-errors:
				logger.WithFields(log.Fields{
					"caller":  "kernel.data.redis.factory",
					"address": k.Config.Cache.Address,
					"detail":  err.Error(),
				}).Error("failed to connect to redis database")
			}
		})
	}

	cleanup := func() {
		if redisPool != nil {
			if err := redisPool.Close(); err != nil {
				logger.WithFields(log.Fields{
					"caller": "kernel.data.redis.factory",
					"detail": err.Error(),
				}).Error("failed to close redis client connection")
			}
		}
	}

	return redisPool, cleanup
}

func getRedisPool(ctx context.Context, k domain.KernelStore) error {
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
		return errPing
	}

	return nil
}
