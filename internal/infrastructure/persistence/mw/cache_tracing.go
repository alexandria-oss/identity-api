package mw

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/tracingutil"
	"go.opencensus.io/trace"
	"strings"
	"time"
)

type CacheTracing struct {
	Next repository.Cache
}

func (c CacheTracing) Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/cache.write")
	span.AddAttributes(trace.StringAttribute("operation", "write"))
	defer c.injectTracing(span, "write to cache", err)

	err = c.Next.Write(ctxT, table, key, value, duration)
	return
}

func (c CacheTracing) Read(ctx context.Context, table, key string) (res string, err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/cache.read")
	span.AddAttributes(trace.StringAttribute("operation", "read"))
	defer c.injectTracing(span, "read from cache", err)

	res, err = c.Next.Read(ctxT, table, key)
	return
}

func (c CacheTracing) Invalidate(ctx context.Context, table, key string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/cache.invalidate")
	span.AddAttributes(trace.StringAttribute("operation", "invalidate"))
	defer c.injectTracing(span, "invalidate data from cache", err)

	err = c.Next.Invalidate(ctxT, table, key)
	return
}

func (c CacheTracing) injectTracing(span *trace.Span, operation string, err error) {
	var status trace.Status
	if err != nil {
		status = trace.Status{
			Code:    tracingutil.CodeFromError(err),
			Message: err.Error(),
		}
	} else {
		status = trace.Status{
			Code:    trace.StatusCodeOK,
			Message: strings.ToLower(operation),
		}
	}

	span.SetStatus(status)
	span.End()
}
