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

package handler

import (
	"encoding/json"
	"github.com/alexandria-oss/common-go/httputil"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

// User HTTP transport.Handler container
type User struct {
	command service.UserCommandHandler
	query   service.UserQuery
}

// NewUser Create a new User HTTP Handler container
func NewUser(cmd service.UserCommandHandler, q service.UserQuery) *User {
	return &User{
		command: cmd,
		query:   q,
	}
}

/* HTTP Handler imp */

func (User) GetName() string {
	return "user"
}

func (u User) SetRoutes(r *mux.Router) {
	r.Path("/user/{key}").Methods(http.MethodGet).HandlerFunc(u.get)
	r.Path("/user").Methods(http.MethodGet).HandlerFunc(u.list)

	r.Path("/user/{username}/enable").Methods(http.MethodPatch).HandlerFunc(u.enable)
	r.Path("/user/{username}/disable").Methods(http.MethodPatch).HandlerFunc(u.disable)
	r.Path("/user/{username}/remove").Methods(http.MethodDelete).HandlerFunc(u.remove)
}

/* Actual Handlers */

func (u User) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var user *aggregate.UserRoot
	var err error
	// Username fetch by default
	if r.URL.Query().Get("by_id") == "true" {
		user, err = u.query.GetByID(r.Context(), mux.Vars(r)["key"])
	} else {
		user, err = u.query.Get(r.Context(), mux.Vars(r)["key"])
	}

	if err != nil {
		httputil.RespondErrorJSON(err, w)
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}

func (u User) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	users, nextToken, err := u.query.List(r.Context(), &domain.Criteria{
		FilterBy: domain.FilterMap{
			"name":        r.URL.Query().Get("name"),
			"email":       r.URL.Query().Get("email"),
			"middle_name": r.URL.Query().Get("middle_name"),
			"family_name": r.URL.Query().Get("family_name"),
			"locale":      r.URL.Query().Get("locale"),
			"disabled":    r.URL.Query().Get("disabled"),
		},
		Token:   domain.PaginationToken(r.URL.Query().Get("page_token")),
		Limit:   domain.NewLimit(r.URL.Query().Get("page_size")),
		OrderBy: domain.Order(r.URL.Query().Get("order_by")),
	})
	if err != nil {
		httputil.RespondErrorJSON(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(struct {
		Users     []*aggregate.UserRoot `json:"users"`
		NextToken string                `json:"next_token"`
	}{
		Users:     users,
		NextToken: string(nextToken),
	})
}

func (u User) enable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := u.command.Enable(command.Enable{Ctx: r.Context(), ID: mux.Vars(r)["username"]}); err != nil {
		httputil.RespondErrorJSON(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(struct{}{})
}

func (u User) disable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := u.command.Disable(command.Disable{
		Ctx: r.Context(),
		ID:  mux.Vars(r)["username"],
	}); err != nil {
		httputil.RespondErrorJSON(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(struct{}{})
}

func (u User) remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := u.command.Remove(command.Remove{
		Ctx: r.Context(),
		ID:  mux.Vars(r)["username"],
	}); err != nil {
		httputil.RespondErrorJSON(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(struct{}{})
}
