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
	"github.com/alexandria-oss/identity-api/pkg/service/wrapper"
	"github.com/alexandria-oss/identity-api/pkg/transport"
	"github.com/alexandria-oss/identity-api/pkg/transport/handler"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
)

var httpSet = wire.NewSet(
	domain.NewKernelStore,
	logging.NewLogger,
	dependency.InjectUserCommandHandler,
	provideUserCommandHandler,
	dependency.InjectUserQuery,
	provideUserQuery,
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

func provideUserCommandHandler(svc *cmdhandler.UserHandlerImp, logger *log.Logger) service.UserCommandHandler {
	return wrapper.NewUserCommandHandler(svc, logger)
}

func provideUserQuery(svc *query.UserQueryImp, logger *log.Logger) service.UserQuery {
	return wrapper.NewUserQuery(svc, logger)
}

func provideHandlers(user *handler.User) []transport.Handler {
	return []transport.Handler{user}
}

func InjectHTTP() (*transport.HTTPServer, func(), error) {
	wire.Build(httpSet)

	return &transport.HTTPServer{}, nil, nil
}
