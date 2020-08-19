package handler

import (
	"encoding/json"
	"github.com/alexandria-oss/common-go/httputil"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/alexandria-oss/identity-api/pkg/transport/observability"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	command service.UserCommandHandler
	query   service.UserQuery
}

func NewUser(cmd service.UserCommandHandler, q service.UserQuery) *User {
	return &User{
		command: cmd,
		query:   q,
	}
}

func (User) GetName() string {
	return "user"
}

func (u User) SetRoutes(r *mux.Router) {
	r.Path("/user/{username}").Methods(http.MethodGet).Handler(observability.TraceHTTP(u.get, true))
	r.Path("/user").Methods(http.MethodGet).Handler(observability.TraceHTTP(u.list, true))
}

func (u User) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	user, err := u.query.Get(r.Context(), mux.Vars(r)["username"])
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
	}

	_ = json.NewEncoder(w).Encode(struct {
		Users     []*aggregate.UserRoot `json:"users"`
		NextToken string                `json:"next_token"`
	}{
		Users:     users,
		NextToken: string(nextToken),
	})
}
