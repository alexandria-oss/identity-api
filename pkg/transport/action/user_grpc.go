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

package action

import (
	"context"
	"github.com/alexandria-oss/common-go/grpcutil"
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/pkg/adapter"
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/alexandria-oss/identity-api/proto"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

// User gRPC transport.GRPCService container
type User struct {
	cmd   service.UserCommandHandler
	query service.UserQuery
}

// NewUser Create a new User gRPC Handler container
func NewUser(cmd service.UserCommandHandler, q service.UserQuery) *User {
	return &User{
		cmd:   cmd,
		query: q,
	}
}

/* gRPC service imp */

func (u User) SetRoutes(s *grpc.Server) {
	proto.RegisterIdentityServer(s, &u)
}

func (User) GetName() string {
	return "user"
}

/* Actual Action(s) */

func (u *User) Enable(ctx context.Context, r *proto.CommandRequest) (*proto.Empty, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.enable")
	defer span.End()

	if err := u.cmd.Enable(command.Enable{
		Ctx: ctx,
		ID:  r.Username,
	}); err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return &proto.Empty{}, nil
}

func (u *User) Disable(ctx context.Context, r *proto.CommandRequest) (*proto.Empty, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.disable")
	defer span.End()

	if err := u.cmd.Disable(command.Disable{
		Ctx: ctx,
		ID:  r.Username,
	}); err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return &proto.Empty{}, nil
}

func (u *User) Remove(ctx context.Context, r *proto.CommandRequest) (*proto.Empty, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.remove")
	defer span.End()

	if err := u.cmd.Remove(command.Remove{
		Ctx: ctx,
		ID:  r.Username,
	}); err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return &proto.Empty{}, nil
}

func (u *User) Get(ctx context.Context, r *proto.GetRequest) (*proto.User, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.get")
	defer span.End()

	user, err := u.query.Get(ctx, r.Username)
	if err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return adapter.UserToGRPC(user), nil
}

func (u *User) GetByID(ctx context.Context, r *proto.GetByIDRequest) (*proto.User, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.get_by_id")
	defer span.End()

	user, err := u.query.GetByID(ctx, r.Id)
	if err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return adapter.UserToGRPC(user), nil
}

func (u *User) List(ctx context.Context, r *proto.ListRequest) (*proto.ListResponse, error) {
	ctx, span := trace.StartSpan(ctx, "identity/grpc.list")
	defer span.End()

	users, nextToken, err := u.query.List(ctx, &domain.Criteria{
		FilterBy: map[string]string{
			"name":        r.Filter["name"],
			"email":       r.Filter["email"],
			"middle_name": r.Filter["middle_name"],
			"family_name": r.Filter["family_name"],
			"locale":      r.Filter["locale"],
			"disabled":    r.Filter["disabled"],
		},
		Token:   domain.PaginationToken(r.Token),
		Limit:   domain.Limit(r.Limit),
		OrderBy: domain.Order(r.OrderBy),
	})
	if err != nil {
		return nil, grpcutil.RespondError(err)
	}

	return &proto.ListResponse{
		Users:         adapter.BulkUserToGRPC(users),
		NextPageToken: nextToken.GetPrimitive(),
	}, nil
}
