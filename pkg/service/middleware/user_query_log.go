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
	"context"
	"fmt"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/pkg/service"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserQueryLog struct {
	Logger *log.Logger
	Next   service.UserQuery
}

func (q UserQueryLog) Get(ctx context.Context, username string) (user *aggregate.UserRoot, err error) {
	defer func(begin time.Time) {
		logQuery(q.Logger, "service.query.get", "get", fmt.Sprintf("username: %s", username),
			user, err, begin)
	}(time.Now())

	user, err = q.Next.Get(ctx, username)
	return
}

func (q UserQueryLog) GetByID(ctx context.Context, id string) (user *aggregate.UserRoot, err error) {
	defer func(begin time.Time) {
		logQuery(q.Logger, "service.query.get_by_id", "get_by_id", fmt.Sprintf("id: %s", id),
			user, err, begin)
	}(time.Now())

	user, err = q.Next.GetByID(ctx, id)
	return
}

func (q UserQueryLog) List(ctx context.Context, criteria *domain.Criteria) (users []*aggregate.UserRoot,
	token domain.PaginationToken, err error) {
	defer func(begin time.Time) {
		logQuery(q.Logger, "service.query.list", "list",
			fmt.Sprintf("token: %v, limit: %v, order_by: %v, filter_by: %+v", criteria.Token, criteria.Limit, criteria.OrderBy,
				criteria.FilterBy),
			token, err, begin)
	}(time.Now())

	users, token, err = q.Next.List(ctx, criteria)
	return
}
