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
	"fmt"
	"github.com/alexandria-oss/identity-api/pkg/transport/resiliency"
	rpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

// GRPCServer gRPC Server proxy
type GRPCServer struct {
	server   *grpc.Server
	ctx      context.Context
	logger   *log.Logger
	services []GRPCService
}

// NewGRPCServer GRCPServer factory method - Create a new preconfigured gRPC Server instance
func NewGRPCServer(ctx context.Context, logger *log.Logger, services ...GRPCService) *GRPCServer {
	srv := &GRPCServer{
		ctx:      ctx,
		logger:   logger,
		services: services,
	}

	srv.server = srv.setInterceptors()
	srv.MapRoutes()

	grpc_prometheus.Register(srv.server)
	return srv
}

// GetServer Obtain GRPCServer's root grpc.Server reference
func (s *GRPCServer) GetServer() *grpc.Server {
	return s.server
}

// AddGRPCServices Runtime GRPCServer service injection using an strategy-like pattern
func (s *GRPCServer) AddGRPCServices(services ...GRPCService) {
	for _, svc := range services {
		s.services = append(s.services, svc)
	}
}

// MapRoutes Runtime GRPCServer service mapping to internal server
func (s GRPCServer) MapRoutes() {
	// Inject OpenCensus distributed tracing in public router
	// Note: This is done here to avoid Health Check and metric endpoints getting traced
	for _, h := range s.services {
		h.SetRoutes(s.server)
		s.logger.WithField("caller", "transport.grpc.routing").
			Info(fmt.Sprintf("service '%s' mapped to public server", h.GetName()))
	}
}

func (s *GRPCServer) setInterceptors() *grpc.Server {
	return grpc.NewServer(
		grpc.StatsHandler(&ocgrpc.ServerHandler{
			IsPublicEndpoint: false,
			StartOptions:     trace.StartOptions{},
		}),
		grpc.UnaryInterceptor(rpcmw.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(s.logger.WithContext(s.ctx)),
			resiliency.GRPCRateLimit(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
}
