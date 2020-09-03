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
	"context"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/dependency"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/logging"
)

// Integration testing
func main() {
	ctx := context.Background()
	logger := logging.NewLogger(domain.NewKernelStore())
	dependency.SetContext(ctx)
	usrQuery, cleanup := dependency.InjectUserQuery()
	defer cleanup()

	logger.Info("loaded sandbox stage")

	user, err := usrQuery.Get(ctx, "drossus")
	if err != nil {
		panic(exception.GetDescription(err))
	}

	logger.Printf("%+v", user.User)

	users, token, err := usrQuery.List(ctx, &domain.Criteria{
		FilterBy: domain.FilterMap{"email": ""},
		Token:    "",
		Limit:    domain.NewLimit("1"),
		OrderBy:  "",
	})
	if err != nil {
		panic(exception.GetDescription(err))
	}

	for _, user := range users {
		logger.Printf("user: %+v", user.User)
	}

	logger.Printf("next_token: %s", token)

	/*
		usrCmd := dependency.InjectUserCommandHandler()
		err = usrCmd.Enable(command.Enable{
			Ctx: ctx,
			ID:  user.Username,
		})
		if err != nil {
			panic(exception.GetDescription(err))
		}

		log.Printf("user %s enable", user.Username)*/
}
