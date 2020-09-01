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
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/pkg/service"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserCommandHandlerLog struct {
	Logger *log.Logger
	Next   service.UserCommandHandler
}

func (l UserCommandHandlerLog) Enable(cmd command.Enable) (err error) {
	defer func(begin time.Time) {
		logCommand(l.Logger, "service.user.enable", "enable", fmt.Sprintf("id: %s", cmd.ID), err, begin)
	}(time.Now())

	err = l.Next.Enable(cmd)
	return
}

func (l UserCommandHandlerLog) Disable(cmd command.Disable) (err error) {
	defer func(begin time.Time) {
		logCommand(l.Logger, "service.user.disable", "disable", fmt.Sprintf("id: %s", cmd.ID), err, begin)
	}(time.Now())

	err = l.Next.Disable(cmd)
	return
}

func (l UserCommandHandlerLog) Remove(cmd command.Remove) (err error) {
	defer func(begin time.Time) {
		logCommand(l.Logger, "service.user.remove", "remove", fmt.Sprintf("id: %s", cmd.ID), err, begin)
	}(time.Now())

	err = l.Next.Remove(cmd)
	return
}
