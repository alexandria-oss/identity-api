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

package main

import (
	"github.com/alexandria-oss/identity-api/internal/infrastructure/dependency"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/logging"
	"github.com/alexandria-oss/identity-api/pkg/dep"
	"github.com/oklog/run"
	"github.com/sethvargo/go-signalcontext"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func main() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()
	logger := logging.NewLogger()

	dep.SetContext(ctx)
	dependency.SetContext(ctx)
	transport, cleanup, err := dep.InjectTransport()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	var g run.Group
	{
		l, err := net.Listen("tcp", transport.HTTP.GetServer().Addr)
		if err != nil {
			panic(err)
		}

		g.Add(func() error {
			return http.Serve(l, transport.HTTP.GetServer().Handler)
		}, func(err error) {
			if err != nil {
				logger.WithFields(log.Fields{
					"caller": "main.http",
					"detail": err.Error(),
				}).Error("http tcp server failed to start")
			}

			_ = l.Close()
		})
	}
	{
		l, err := net.Listen("tcp", transport.GetKernel().Config.Transport.GRPC.Address)
		if err != nil {
			panic(err)
		}

		g.Add(func() error {
			return transport.GRPC.GetServer().Serve(l)
		}, func(err error) {
			if err != nil {
				logger.WithFields(log.Fields{
					"caller": "main.grpc",
					"detail": err.Error(),
				}).Error("grpc tcp server failed to start")
			}

			_ = l.Close()
		})
	}
	{
		g.Add(func() error {
			select {
			case <-ctx.Done():
				return nil
			}
		}, func(_ error) {
			return
		})
	}

	if err := g.Run(); err != nil {
		logger.WithFields(log.Fields{
			"caller": "main",
			"detail": err.Error(),
		}).Error("goroutines failed to end gracefully")
	}
}
