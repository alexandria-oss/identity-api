package persistence

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

type CacheRedis struct {
	db *redis.Client
	mu *sync.RWMutex
}

func NewCacheRedis(pool *redis.Client) *CacheRedis {
	return &CacheRedis{
		db: pool,
		mu: new(sync.RWMutex),
	}
}

func (c *CacheRedis) Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if conn := c.db.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()

		return conn.Set(ctx, c.generateComposedKey(table, key), value, duration).Err()
	}

	return exception.NewNetworkCall("redis", c.db.Options().Addr)
}

func (c *CacheRedis) Read(ctx context.Context, table, key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if conn := c.db.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()

		res, err := conn.Get(ctx, c.generateComposedKey(table, key)).Result()
		if err != nil {
			log.Print(err)
		}

		return res, err
	}

	return "", exception.NewNetworkCall("redis", c.db.Options().Addr)
}

func (c *CacheRedis) Invalidate(ctx context.Context, table, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if conn := c.db.Conn(ctx); conn != nil {
		defer func() {
			_ = conn.Close()
		}()

		return conn.Del(ctx, c.generateComposedKey(table, key)).Err()
	}

	return exception.NewNetworkCall("redis", c.db.Options().Addr)
}

func (c CacheRedis) generateComposedKey(primary, secondary string) string {
	return fmt.Sprintf("%s:%s", primary, secondary)
}
