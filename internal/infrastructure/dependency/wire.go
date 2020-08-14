// +build wireinject

package dependency

import (
	"github.com/alexandria-oss/identity-api/internal/application/command/cmdhandler"
	"github.com/alexandria-oss/identity-api/internal/application/query"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/repository"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/driver"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/persistence"
	"github.com/google/wire"
)

var userQuery = wire.NewSet(
	domain.NewKernelStore,
	wire.Bind(new(repository.User), new(*persistence.UserAWSRepository)),
	driver.NewCognitoSession,
	persistence.NewUserAWSRepository,
	query.NewUserQuery,
)

func InjectUserQuery() *query.UserQueryImp {
	wire.Build(userQuery)
	return &query.UserQueryImp{}
}

func InjectUserCommandHandler() *cmdhandler.UserImp {
	wire.Build(
		domain.NewKernelStore,
		wire.Bind(new(repository.User), new(*persistence.UserAWSRepository)),
		driver.NewCognitoSession,
		persistence.NewUserAWSRepository,
		cmdhandler.NewUserCommandHandler,
	)
	return &cmdhandler.UserImp{}
}
