// +build wireinject

package dependency

import (
	"github.com/alexandria-oss/identity-api/internal/command"
	"github.com/alexandria-oss/identity-api/internal/common"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/infrastructure"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/driver"
	"github.com/alexandria-oss/identity-api/internal/query"
	"github.com/google/wire"
)

var userQuery = wire.NewSet(
	common.NewKernelStore,
	wire.Bind(new(domain.UserQueryRepository), new(*infrastructure.UserQueryAWSRepository)),
	driver.NewCognitoSession,
	infrastructure.NewUserQueryAWSRepository,
	query.NewUserQueryImp,
)

func InjectUserQuery() *query.UserQueryImp {
	wire.Build(userQuery)
	return &query.UserQueryImp{}
}

func InjectUserCommand() *command.UserCommandImp {
	wire.Build(
		common.NewKernelStore,
		wire.Bind(new(domain.UserCommandRepository), new(*infrastructure.UserCommandAWSRepository)),
		driver.NewCognitoSession,
		infrastructure.NewUserCommandAWSRepository,
		command.NewUserCommand,
	)
	return &command.UserCommandImp{}
}
