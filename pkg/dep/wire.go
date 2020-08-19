// +build wireinject

package dep

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/application/command/cmdhandler"
	"github.com/alexandria-oss/identity-api/internal/application/query"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/dependency"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/logging"
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/alexandria-oss/identity-api/pkg/transport"
	"github.com/alexandria-oss/identity-api/pkg/transport/handler"
	"github.com/google/wire"
)

var httpSet = wire.NewSet(
	domain.NewKernelStore,
	logging.NewLogger,
	wire.Bind(new(service.UserCommandHandler), new(*cmdhandler.UserHandlerImp)),
	dependency.InjectUserCommandHandler,
	wire.Bind(new(service.UserQuery), new(*query.UserQueryImp)),
	dependency.InjectUserQuery,
	handler.NewUser,
	provideHandlers,
	transport.NewHTTPServer,
)

var ctx = context.Background()

func SetContext(parentCtx context.Context) {
	ctx = parentCtx
}

func provideContext() context.Context {
	return ctx
}

func provideHandlers(user *handler.User) []transport.Handler {
	return []transport.Handler{user}
}

func InjectHTTP() (*transport.HTTPServer, func(), error) {
	wire.Build(httpSet)

	return &transport.HTTPServer{}, nil, nil
}
