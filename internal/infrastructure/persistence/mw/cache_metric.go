package mw

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type CacheMetric struct {
	Next repository.Cache
}

var (
	usageGaugeCache  prometheus.Gauge
	errCounterCache  *prometheus.CounterVec
	hitCounterCache  *prometheus.CounterVec
	hitDurationCache *prometheus.SummaryVec

	metricSingletonCache = new(sync.Once)
)

func NewCacheMetric(n repository.Cache, l *log.Logger) CacheMetric {
	metricSingletonCache.Do(func() {
		labels := []string{"caller", "error"}
		usageGaugeCache = prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_cache_usage",
			Help:        "user cache global concurrent usage",
			ConstLabels: nil,
		})
		errCounterCache = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_cache_error_count",
			Help:        "user cache total errors occurred",
			ConstLabels: nil,
		}, labels)
		hitCounterCache = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_cache_hit_count",
			Help:        "user cache total hits",
			ConstLabels: nil,
		}, labels)
		hitDurationCache = prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:   "alexandria",
			Subsystem:   "identity",
			Name:        "user_cache_hit_duration",
			Help:        "user cache latency by hit in microseconds (us)",
			ConstLabels: nil,
			Objectives:  nil,
			MaxAge:      0,
			AgeBuckets:  0,
			BufCap:      0,
		}, labels)
		registerMetricsCache(l)
	})

	return CacheMetric{
		Next: n,
	}
}

func registerMetricsCache(l *log.Logger) {
	failedMetricRegistry("user_cache_usage", prometheus.Register(usageGaugeCache), l)
	failedMetricRegistry("user_cache_error_count", prometheus.Register(errCounterCache), l)
	failedMetricRegistry("user_cache_hit_count", prometheus.Register(hitCounterCache), l)
	failedMetricRegistry("user_cache_hit_duration", prometheus.Register(hitDurationCache), l)
}

func (c CacheMetric) Write(ctx context.Context, table, key string, value interface{}, duration time.Duration) (err error) {
	begin := time.Now()
	usageGaugeCache.Inc()

	err = c.Next.Write(ctx, table, key, value, duration)
	defer c.injectMetrics("write", begin, err)
	return
}

func (c CacheMetric) Read(ctx context.Context, table, key string) (res string, err error) {
	begin := time.Now()
	usageGaugeCache.Inc()

	res, err = c.Next.Read(ctx, table, key)
	defer c.injectMetrics("read", begin, err)
	return
}

func (c CacheMetric) Invalidate(ctx context.Context, table, key string) (err error) {
	begin := time.Now()
	usageGaugeCache.Inc()

	err = c.Next.Invalidate(ctx, table, key)
	defer c.injectMetrics("invalidate", begin, err)
	return
}

func (u CacheMetric) injectMetrics(caller string, begin time.Time, err error) {
	lvs := prometheus.Labels{
		"caller": "cache." + caller,
		"error":  fmt.Sprint(err != nil),
	}
	if err != nil {
		errCounterCache.With(lvs).Inc()
	}
	hitCounterCache.With(lvs).Inc()
	hitDurationCache.With(lvs).Observe(time.Since(begin).Seconds())
	usageGaugeCache.Dec()
}
