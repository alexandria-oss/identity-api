package mw

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type UserRepositoryMetric struct {
	Next repository.User
}

var (
	usageGauge  prometheus.Gauge
	errCounter  *prometheus.CounterVec
	hitCounter  *prometheus.CounterVec
	hitDuration *prometheus.SummaryVec

	metricSingleton = new(sync.Once)
)

func init() {
	metricSingleton.Do(func() {
		labels := []string{"caller", "error"}
		usageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_repository_usage",
			Help:        "user repository global concurrent usage",
			ConstLabels: nil,
		})
		errCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_repository_error_count",
			Help:        "user repository total errors occurred by command or query",
			ConstLabels: nil,
		}, labels)
		hitCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_repository_hit_count",
			Help:        "user repository total hits by command or query",
			ConstLabels: nil,
		}, labels)
		hitDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_repository_hit_duration",
			Help:        "user repository latency by hit in microseconds (us)",
			ConstLabels: nil,
			Objectives:  nil,
			MaxAge:      0,
			AgeBuckets:  0,
			BufCap:      0,
		}, labels)
	})
}

func (u UserRepositoryMetric) Remove(ctx context.Context, id string) (err error) {
	usageGauge.Inc()
	defer u.injectMetrics("remove", time.Now(), err)
	err = u.Next.Remove(ctx, id)
	return
}

func (u UserRepositoryMetric) Restore(ctx context.Context, id string) (err error) {
	usageGauge.Inc()
	defer u.injectMetrics("restore", time.Now(), err)
	err = u.Next.Restore(ctx, id)
	return
}

func (u UserRepositoryMetric) HardRemove(ctx context.Context, id string) (err error) {
	usageGauge.Inc()
	defer u.injectMetrics("hardRemove", time.Now(), err)
	err = u.Next.Remove(ctx, id)
	return
}

func (u UserRepositoryMetric) FetchOne(ctx context.Context, byUsername bool, key string) (user *aggregate.UserRoot, err error) {
	usageGauge.Inc()
	defer u.injectMetrics("fetchOne", time.Now(), err)
	user, err = u.Next.FetchOne(ctx, byUsername, key)
	return
}

func (u UserRepositoryMetric) Fetch(ctx context.Context, criteria domain.Criteria) (users []*aggregate.UserRoot, err error) {
	usageGauge.Inc()
	defer u.injectMetrics("fetch", time.Now(), err)
	users, err = u.Next.Fetch(ctx, criteria)
	return
}

func (u UserRepositoryMetric) injectMetrics(caller string, begin time.Time, err error) {
	lvs := prometheus.Labels{
		"caller": "repository." + caller,
		"error":  fmt.Sprint(err != nil),
	}
	if err != nil {
		errCounter.With(lvs).Inc()
	}
	hitCounter.With(lvs).Inc()
	hitDuration.With(lvs).Observe(time.Since(begin).Seconds())
	usageGauge.Dec()
}
