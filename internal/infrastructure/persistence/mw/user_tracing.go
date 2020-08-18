package mw

import (
	"context"
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
	defer func() {
		var status trace.Status
		if err != nil {
			status = trace.Status{
				Code:    tracingutil.CodeFromError(err),
				Message: err.Error(),
			}
		} else {
			status = trace.Status{
				Code:    trace.StatusCodeOK,
				Message: "removed row",
			}
		}

		span.SetStatus(status)
		span.End()
	}()

	err = u.Next.Remove(ctxT, id)
	return
}

func (u UserRepositoryTracing) Restore(ctx context.Context, id string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.restore")
	span.AddAttributes(trace.StringAttribute("operation", "restore"))
	defer func() {
		var status trace.Status
		if err != nil {
			status = trace.Status{
				Code:    tracingutil.CodeFromError(err),
				Message: err.Error(),
			}
		} else {
			status = trace.Status{
				Code:    trace.StatusCodeOK,
				Message: "restored row",
			}
		}

		span.SetStatus(status)
		span.End()
	}()

	err = u.Next.Restore(ctxT, id)
	return
}

func (u UserRepositoryTracing) HardRemove(ctx context.Context, id string) (err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.hardRemove")
	span.AddAttributes(trace.StringAttribute("operation", "hardRemove"))
	defer func() {
		var status trace.Status
		if err != nil {
			status = trace.Status{
				Code:    tracingutil.CodeFromError(err),
				Message: err.Error(),
			}
		} else {
			status = trace.Status{
				Code:    trace.StatusCodeOK,
				Message: "permanently removed row",
			}
		}

		span.SetStatus(status)
		span.End()
	}()

	err = u.Next.HardRemove(ctxT, id)
	return
}

func (u UserRepositoryTracing) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.fetchOne")
	span.AddAttributes(trace.StringAttribute("operation", "fetchOne"))
	defer func() {
		var status trace.Status
		if err != nil {
			status = trace.Status{
				Code:    tracingutil.CodeFromError(err),
				Message: err.Error(),
			}
		} else {
			status = trace.Status{
				Code:    trace.StatusCodeOK,
				Message: "fetched row",
			}
		}

		span.SetStatus(status)
		span.End()
	}()

	user, err = u.Next.FetchOne(ctxT, byUsername, key)
	return
}

func (u UserRepositoryTracing) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot, err error) {
	ctxT, span := trace.StartSpan(ctx, "identity/repository.fetch")
	span.AddAttributes(trace.StringAttribute("operation", "fetch"))
	defer func() {
		var status trace.Status
		if err != nil {
			status = trace.Status{
				Code:    tracingutil.CodeFromError(err),
				Message: err.Error(),
			}
		} else {
			status = trace.Status{
				Code:    trace.StatusCodeOK,
				Message: "fetched rows",
			}
		}

		span.SetStatus(status)
		span.End()
	}()

	users, err = u.Next.Fetch(ctxT, criteria)
	return
}
