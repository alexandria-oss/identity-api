package main

import (
	"context"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/dependency"
	"github.com/alexandria-oss/identity-api/internal/infrastructure/logging"
)

func main() {
	ctx := context.Background()
	logger := logging.NewLogger()
	dependency.SetContext(ctx)
	usrQuery, cleanup := dependency.InjectUserQuery()
	defer cleanup()

	logger.Info("loaded sandbox stage")

	user, err := usrQuery.Get(ctx, "drossus")
	if err != nil {
		panic(exception.GetDescription(err))
	}

	logger.Printf("%+v", user.Root)

	users, token, err := usrQuery.List(ctx, &domain.Criteria{
		FilterBy: domain.FilterMap{"email": ""},
		Token:    "",
		Limit:    domain.NewLimit("1"),
		OrderBy:  "",
	})
	if err != nil {
		panic(exception.GetDescription(err))
	}

	for _, user := range users {
		logger.Printf("user: %+v", user.Root)
	}

	logger.Printf("next_token: %s", token)

	/*
		usrCmd := dependency.InjectUserCommandHandler()
		err = usrCmd.Enable(command.Enable{
			Ctx: ctx,
			ID:  user.Username,
		})
		if err != nil {
			panic(exception.GetDescription(err))
		}

		log.Printf("user %s enable", user.Username)*/
}
