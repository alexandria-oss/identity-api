// Copyright 2020 The Alexandria Foundation
//
// Licensed under the GNU Affero General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/pkg/transport/observability"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// TransportFacade Encapsulates transport-tier complexity
type TransportFacade struct {
	kernel domain.KernelStore
	logger *log.Logger
	HTTP   *HTTPServer
	GRPC   *GRPCServer
	// TODO: Add PubSub/AMQP
}

// NewTransportFacade create a new TransportFacade encapsulated sub-ecosystem
func NewTransportFacade(ctx context.Context, kernel domain.KernelStore, logger *log.Logger, handlers []Handler,
	rpcServices []GRPCService) (*TransportFacade, func(), error) {
	// Avoid any operation(s) if nil reference was found
	if ctx == nil || logger == nil || &kernel == nil {
		return nil, nil, errors.New("missing required transport parameters")
	}

	f := &TransportFacade{
		kernel: kernel,
		logger: logger,
		HTTP:   NewHTTPServer(kernel, logger, handlers...),
		GRPC:   NewGRPCServer(ctx, logger, rpcServices...),
	}

	// Enable Observability
	errTracing, cleanup := f.injectTracing()
	if errTracing != nil {
		return nil, nil, errTracing
	}
	if err := f.injectMetrics(); err != nil {
		return nil, cleanup, err
	}
	f.injectHealthCheck()

	return f, cleanup, nil
}

// GetKernel get transport current kernel
func (f TransportFacade) GetKernel() domain.KernelStore {
	return f.kernel
}

func (f TransportFacade) injectMetrics() error {
	pe, err := observability.StartPrometheusMonitoring(f.kernel)
	if err != nil {
		f.logger.WithFields(log.Fields{
			"caller": "transport.facade.metrics",
			"detail": err.Error(),
		}).Error("cannot start prometheus metrics exporter")
		return err
	}

	f.HTTP.GetRouter().Path("/metrics").Handler(pe)
	return nil
}

func (f TransportFacade) injectTracing() (error, func()) {
	je, err := observability.StartJaegerTracing(f.kernel)
	if err != nil {
		f.logger.WithFields(log.Fields{
			"caller": "transport.facade.tracing",
			"detail": err.Error(),
		}).Error("cannot start jaeger agent and collector exporter")
		return err, nil
	}

	return nil, je.Flush
}

func (f TransportFacade) injectHealthCheck() {
	f.HTTP.GetRouter().Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "UP",
		})
	})
}
