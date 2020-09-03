package observability

import (
	ocprom "contrib.go.opencensus.io/exporter/prometheus"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"strings"
	"time"
)

// StartPrometheusMonitoring register required OpenCensus views and register a new prometheus exporter
func StartPrometheusMonitoring(k domain.KernelStore) (*ocprom.Exporter, error) {
	view.SetReportingPeriod(time.Second * 60)

	if err := registerHTTPPrometheus(); err != nil {
		return nil, err
	}

	if err := registerGRPCPrometheus(); err != nil {
		return nil, err
	}

	pe, err := ocprom.NewExporter(ocprom.Options{
		Namespace:   strings.ToLower(k.Service),
		Registry:    nil,
		Registerer:  prometheus.DefaultRegisterer,
		Gatherer:    prometheus.DefaultGatherer,
		OnError:     nil,
		ConstLabels: nil,
	})
	if err != nil {
		return nil, err
	}

	view.RegisterExporter(pe)
	return pe, nil
}

func registerHTTPPrometheus() error {
	if err := view.Register(
		ochttp.ServerLatencyView,
		ochttp.ServerRequestBytesView,
		ochttp.ServerRequestCountByMethod,
		ochttp.ServerRequestCountView,
		ochttp.ServerResponseBytesView,
		ochttp.ServerResponseCountByStatusCode,
	); err != nil {
		return err
	}

	return nil
}

func registerGRPCPrometheus() error {
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		return err
	}

	return nil
}
