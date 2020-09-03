package mw

import (
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/eapache/go-resiliency/retrier"
	log "github.com/sirupsen/logrus"
	"time"
)

type CacheRetryPolicy struct {
	Logger *log.Logger
	Next   repository.Cache
}

func (c CacheRetryPolicy) Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) error {
	// Circuit Breaker
	// Note: CB Timeout policy must be in ms and should be equal to retry policy's total time
	// CB Timeout policy f(x) = x(t)(1000)   // where x = retry times - t = time in sec - *result units in ms
	hystrix.DefaultTimeout = int(time.Millisecond * 30000)
	return hystrix.DoC(ctx, "cache_write", func(ctxCB context.Context) error {
		// Retry Policy
		// Actual policy: Try 3 times every 10 sec
		r := retrier.New(retrier.ConstantBackoff(3, time.Second*10), retrier.WhitelistClassifier{
			exception.NetworkCall,
		})
		err := r.RunCtx(ctxCB, func(ctxR context.Context) error {
			if err := c.Next.Write(ctxR, table, key, value, duration); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			if errors.Is(err, exception.NetworkCall) {
				c.Logger.WithFields(log.Fields{
					"caller": "kernel.data.cache.write",
					"detail": err.Error(),
				}).Error("failed to connect to cache database")
				return errors.New("could not connect to cache database")
			}

			return err
		}

		return nil
	}, func(_ context.Context, err error) error {
		return err
	})
}

func (c CacheRetryPolicy) Read(ctx context.Context, table, key string) (string, error) {
	output := make(chan string)
	// Circuit Breaker
	// Note: CB Timeout policy must be in ms and should be equal to retry policy's total time
	// CB Timeout policy f(x) = x(t)(1000)   // where x = retry times - t = time in sec - *result units in ms
	hystrix.DefaultTimeout = int(time.Millisecond * 30000)
	errs := hystrix.GoC(ctx, "cache_read", func(ctxCB context.Context) error {
		// Retry Policy
		// Actual policy: Try 3 times every 10 sec
		r := retrier.New(retrier.ConstantBackoff(3, time.Second*10), retrier.WhitelistClassifier{
			exception.NetworkCall,
		})
		err := r.RunCtx(ctxCB, func(ctxR context.Context) error {
			res, err := c.Next.Read(ctxR, table, key)
			if err != nil {
				return err
			}

			output <- res
			return nil
		})
		if err != nil {
			if errors.Is(err, exception.NetworkCall) {
				c.Logger.WithFields(log.Fields{
					"caller": "kernel.data.cache.read",
					"detail": err.Error(),
				}).Error("failed to connect to cache database")
				return errors.New("could not connect to cache database")
			}

			return err
		}

		return nil
	}, func(_ context.Context, err error) error {
		return err
	})

	select {
	case res := <-output:
		return res, nil
	case err := <-errs:
		return "", err
	}
}

func (c CacheRetryPolicy) Invalidate(ctx context.Context, table, key string) error {
	// Circuit Breaker
	// Note: CB Timeout policy must be in ms and should be equal to retry policy's total time
	// CB Timeout policy f(x) = x(t)(1000)   // where x = retry times - t = time in sec - *result units in ms
	hystrix.DefaultTimeout = int(time.Millisecond * 30000)
	return hystrix.DoC(ctx, "cache_invalidate", func(ctxCB context.Context) error {
		// Retry Policy
		// Actual policy: Try 3 times every 10 sec
		r := retrier.New(retrier.ConstantBackoff(3, time.Second*10), retrier.WhitelistClassifier{
			exception.NetworkCall,
		})
		err := r.RunCtx(ctxCB, func(ctxR context.Context) error {
			if err := c.Next.Invalidate(ctxR, table, key); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			if errors.Is(err, exception.NetworkCall) {
				c.Logger.WithFields(log.Fields{
					"caller": "kernel.data.cache.invalidate",
					"detail": err.Error(),
				}).Error("failed to connect to cache database")
				return errors.New("could not connect to cache database")
			}

			return err
		}

		return nil
	}, func(_ context.Context, err error) error {
		return err
	})
}
