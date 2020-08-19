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
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/pkg/transport/observability"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// HTTPServer HTTP Server facade
type HTTPServer struct {
	server   *http.Server
	kernel   domain.KernelStore
	logger   *log.Logger
	router   *mux.Router
	handlers []Handler
}

// NewHTTPServer HTTPServer factory method - Create a new preconfigured HTTP Server instance
func NewHTTPServer(k domain.KernelStore, l *log.Logger, handlers ...Handler) (*HTTPServer, error) {
	srv := &HTTPServer{
		server: nil,
		kernel: k,
		logger: l,
		router: mux.NewRouter(),
	}

	srv.setMiddlewares()

	if err := srv.injectMetrics(); err != nil {
		return nil, err
	}

	if err := srv.injectTracing(); err != nil {
		return nil, err
	}

	srv.AddHandlers(handlers...)
	srv.MapRoutes()

	srv.server = &http.Server{
		Addr:              k.Config.Transport.HTTP.Address,
		Handler:           muxhandlers.CombinedLoggingHandler(srv.logger.Writer(), srv.router),
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	return srv, nil
}

// GetServer Obtain HTTPServer's root http.Server reference
func (s *HTTPServer) GetServer() *http.Server {
	return s.server
}

// AddHandlers Runtime HTTPServer handler injection using an strategy-like pattern
func (s *HTTPServer) AddHandlers(hs ...Handler) {
	for _, h := range hs {
		s.handlers = append(s.handlers, h)
	}
}

// MapRoutes Runtime HTTPServer handler mapping to internal mux
func (s *HTTPServer) MapRoutes() {
	public := s.router.PathPrefix("/" + s.kernel.APIVersion).Subrouter()
	for _, h := range s.handlers {
		h.SetRoutes(public)
		s.logger.WithField("caller", "transport.http.routing").
			Info(fmt.Sprintf("route '/%s' mapped to public router ('/%s')", h.GetName(), s.kernel.APIVersion))
	}
}

func (s HTTPServer) setMiddlewares() {
	s.router.Use(muxhandlers.RecoveryHandler())
	s.router.Use(muxhandlers.CORS(
		muxhandlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		}),
		muxhandlers.AllowedOrigins([]string{"*"}),
	))
	s.router.Use(muxhandlers.CompressHandler)
}

func (s HTTPServer) injectMetrics() error {
	pe, err := observability.InjectPrometheusHTTP(s.kernel)
	if err != nil {
		s.logger.WithFields(log.Fields{
			"caller": "transport.http.metrics",
			"detail": err.Error(),
		}).Error("cannot start prometheus metrics exporter")
		return err
	}
	s.router.Path("/metrics").Handler(pe)
	return nil
}

func (s HTTPServer) injectTracing() error {
	_, err := observability.InjectJaegerHTTP(s.kernel)
	if err != nil {
		s.logger.WithFields(log.Fields{
			"caller": "transport.http.tracing",
			"detail": err.Error(),
		}).Error("cannot start jaeger agent and collector exporter")
		return err
	}
	return nil
}
