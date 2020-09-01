// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func InjectHTTP() (*transport.HTTPServer, func(), error) {
	kernelStore := domain.NewKernelStore()
	logger := logging.NewLogger()
	userHandlerImp, cleanup := dependency.InjectUserCommandHandler()
	userCommandHandler := provideUserCommandHandler(userHandlerImp, logger)
	userQueryImp, cleanup2 := dependency.InjectUserQuery()
	userQuery := provideUserQuery(userQueryImp, logger)
	user := handler.NewUser(userCommandHandler, userQuery)
	v := provideHandlers(user)
	httpServer, err := transport.NewHTTPServer(kernelStore, logger, v...)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return httpServer, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var httpSet = wire.NewSet(domain.NewKernelStore, logging.NewLogger, dependency.InjectUserCommandHandler, provideUserCommandHandler, dependency.InjectUserQuery, provideUserQuery, handler.NewUser, provideHandlers, transport.NewHTTPServer)

var ctx = context.Background()

func SetContext(parentCtx context.Context) {
	ctx = parentCtx
}

func provideContext() context.Context {
	return ctx
}

func provideUserCommandHandler(svc *cmdhandler.UserHandlerImp, logger *logrus.Logger) service.UserCommandHandler {
	return wrapper.NewUserCommandHandler(svc, logger)
}

func provideUserQuery(svc *query.UserQueryImp, logger *logrus.Logger) service.UserQuery {
	return wrapper.NewUserQuery(svc, logger)
}

func provideHandlers(user *handler.User) []transport.Handler {
	return []transport.Handler{user}
}
