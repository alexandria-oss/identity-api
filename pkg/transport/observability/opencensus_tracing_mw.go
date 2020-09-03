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
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
)

func TraceHTTP() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &ochttp.Handler{
			Propagation:      nil,
			Handler:          h,
			StartOptions:     trace.StartOptions{},
			GetStartOptions:  nil,
			IsPublicEndpoint: true,
			FormatSpanName:   nil,
			IsHealthEndpoint: nil,
		}
	}
}
