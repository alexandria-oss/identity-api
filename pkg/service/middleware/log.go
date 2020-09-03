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

package middleware

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func logCommand(l *log.Logger, caller, command, args string, err error, begin time.Time) {
	if err != nil {
		l.WithFields(log.Fields{
			"caller":  caller,
			"command": command,
			"args":    args,
			"took_ms": time.Since(begin).Milliseconds(),
			"detail":  err.Error(),
		}).Error("failed command")
		return
	}

	l.WithFields(log.Fields{
		"caller":  caller,
		"command": command,
		"args":    args,
		"took_ms": time.Since(begin).Milliseconds(),
	}).Info("called command")
}

func logQuery(l *log.Logger, caller, query, args string, output interface{}, err error, begin time.Time) {
	if err != nil {
		l.WithFields(log.Fields{
			"caller":  caller,
			"query":   query,
			"args":    args,
			"took_ms": time.Since(begin).Milliseconds(),
			"detail":  err.Error(),
		}).Error("failed query")
		return
	}

	l.WithFields(log.Fields{
		"caller":  caller,
		"query":   query,
		"args":    args,
		"took_ms": time.Since(begin).Milliseconds(),
		"output":  output,
	}).Info("called query")
}
