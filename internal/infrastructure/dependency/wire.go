// +build wireinject

package dependency

import (
	"github.com/alexandria-oss/identity-api/internal/application/command"
	"github.com/alexandria-oss/identity-api/internal/application/query"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/user"
	"github.com/alexandria-oss/identity-api/internal/infrastructure"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/driver"
	"github.com/google/wire"
)

var userQuery = wire.NewSet(
	domain.NewKernelStore,
	wire.Bind(new(user.UserQueryRepository), new(*infrastructure.UserQueryAWSRepository)),
	driver.NewCognitoSession,
	infrastructure.NewUserQueryAWSRepository,
	query.NewUserQuery,
)

func InjectUserQuery() *query.UserQueryImp {
	wire.Build(userQuery)
	return &query.UserQueryImp{}
}

func InjectUserCommand() *command.UserCommandImp {
	wire.Build(
		domain.NewKernelStore,
		wire.Bind(new(user.UserCommandRepository), new(*infrastructure.UserCommandAWSRepository)),
		driver.NewCognitoSession,
		infrastructure.NewUserCommandAWSRepository,
		command.NewUserCommand,
	)
	return &command.UserCommandImp{}
}
