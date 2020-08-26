// +build wireinject

package dependency

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/application/command/cmdhandler"
	"github.com/alexandria-oss/identity-api/internal/application/query"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/driver"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/logging"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/persistence"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/persistence/mw"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

var dataSet = wire.NewSet(
	domain.NewKernelStore,
	logging.NewLogger,
	provideContext,
	driver.NewRedisClientPool,
	persistence.NewCacheRedis,
	provideCacheRepository,
	driver.NewCognitoSession,
	persistence.NewUserAWSRepository,
	provideUserRepository,
)

func SetContext(parentCtx context.Context) {
	ctx = parentCtx
}

func provideContext() context.Context {
	return ctx
}

func provideCacheRepository(r *persistence.CacheRedis, logger *log.Logger) repository.Cache {
	return mw.WrapCacheRepository(r, logger)
}

func provideUserRepository(r *persistence.UserAWSRepository, c repository.Cache, l *log.Logger) repository.User {
	return mw.WrapUserRepository(r, c, l)
}

func InjectUserQuery() (*query.UserQueryImp, func()) {
	wire.Build(dataSet, query.NewUserQuery)
	return &query.UserQueryImp{}, func() {}
}

func InjectUserCommandHandler() (*cmdhandler.UserHandlerImp, func()) {
	wire.Build(dataSet, cmdhandler.NewUserCommandHandler)
	return &cmdhandler.UserHandlerImp{}, func() {}
}
