package mw

import (
	"context"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/tracingutil"
	"go.opencensus.io/trace"
)

type UserRepositoryTracing struct {
	Next repository.User
}

func (u UserRepositoryTracing) Remove(ctx context.Context, id string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.remove")
	span.AddAttributes(trace.StringAttribute("operation", "remove"))
	defer u.injectTracing(span, "removed user", err)

	err = u.Next.Remove(ctxT, id)
	return
}

func (u UserRepositoryTracing) Restore(ctx context.Context, id string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.restore")
	span.AddAttributes(trace.StringAttribute("operation", "restore"))
	defer u.injectTracing(span, "restored user", err)

	err = u.Next.Restore(ctxT, id)
	return
}

func (u UserRepositoryTracing) HardRemove(ctx context.Context, id string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.hardRemove")
	span.AddAttributes(trace.StringAttribute("operation", "hardRemove"))
	defer u.injectTracing(span, "permanently removed user", err)

	err = u.Next.HardRemove(ctxT, id)
	return
}

func (u UserRepositoryTracing) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.fetchOne")
	span.AddAttributes(trace.StringAttribute("operation", "fetchOne"))
	defer u.injectTracing(span, "fetched user", err)

	user, err = u.Next.FetchOne(ctxT, byUsername, key)
	return
}

func (u UserRepositoryTracing) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot,
	nextToken domain.PaginationToken, err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.fetch")
	span.AddAttributes(trace.StringAttribute("operation", "fetch"))
	defer u.injectTracing(span, "fetched users", err)

	users, nextToken, err = u.Next.Fetch(ctxT, criteria)
	return
}

func (u UserRepositoryTracing) injectTracing(span *trace.Span, operation string, err error) {
	var status trace.Status
	if err != nil {
		status = trace.Status{
			Code:    tracingutil.CodeFromError(err),
			Message: exception.GetDescription(err),
		}
	} else {
		status = trace.Status{
			Code:    trace.StatusCodeOK,
			Message: operation,
		}
	}

	span.SetStatus(status)
	span.End()
}
