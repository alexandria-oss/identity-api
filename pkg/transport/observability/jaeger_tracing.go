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

package observability

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"go.opencensus.io/trace"
	"strings"
)

// StartJaegerTracing register a new Jaeger trace collector
func StartJaegerTracing(k domain.KernelStore) (*jaeger.Exporter, error) {
	je, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: k.Config.Tracing.Collector,
		AgentEndpoint:     k.Config.Tracing.Agent,
		OnError:           nil,
		Username:          k.Config.Tracing.Username,
		Password:          k.Config.Tracing.Password,
		Process: jaeger.Process{
			ServiceName: strings.ToLower(k.Service),
			Tags:        nil,
		},
		BufferMaxCount: 0,
	})

	if err != nil {
		return nil, err
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		// Use probabilistic 1000.0 in prod
		DefaultSampler:             trace.AlwaysSample(),
		IDGenerator:                nil,
		MaxAnnotationEventsPerSpan: 0,
		MaxMessageEventsPerSpan:    0,
		MaxAttributesPerSpan:       0,
		MaxLinksPerSpan:            0,
	})
	return je, nil
}
