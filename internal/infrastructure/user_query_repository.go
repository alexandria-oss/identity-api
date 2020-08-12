package infrastructure

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/common"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"sync"
)

type UserQueryAWSRepository struct {
	client *cognito.CognitoIdentityProvider
	kernel common.KernelStore
	mu     *sync.RWMutex
}

// Factory Method
func NewUserQueryAWSRepository(k common.KernelStore) *UserQueryAWSRepository {
	return &UserQueryAWSRepository{
		client: cognito.New(session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))),
		kernel: k,
		mu:     new(sync.RWMutex),
	}
}

func (r *UserQueryAWSRepository) FetchOne(ctx context.Context, key string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	o, err := r.client.ListUsersWithContext(ctx, &cognito.ListUsersInput{
		AttributesToGet: nil,
		Filter:          aws.String(fmt.Sprintf("sub = \"%s\"", key)),
		Limit:           aws.Int64(1),
		PaginationToken: nil,
		UserPoolId:      aws.String(r.kernel.Config.Cognito.UserPoolID),
	})

	if err != nil {
		return nil, err
	} else if len(o.Users) == 0 {
		return nil, exception.NewCustomError(exception.NotFound, "user")
	}

	return mapToUser(o.Users[0]), nil
}

func (r *UserQueryAWSRepository) Fetch(ctx context.Context, token string, size int, filterMap common.FilterMap) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var statement *string = nil

forLoop:
	for k, v := range filterMap {
		switch k {
		case "name":
			statement = aws.String(fmt.Sprintf("name ^= \"%s\"", v))
			break forLoop
		case "email":
			statement = aws.String(fmt.Sprintf("email ^= \"%s\"", v))
			break forLoop
		case "middle_name":
			statement = aws.String(fmt.Sprintf("middle_name ^= \"%s\"", v))
			break forLoop
		case "family_name":
			statement = aws.String(fmt.Sprintf("family_name ^= \"%s\"", v))
			break forLoop
		case "locale":
			statement = aws.String(fmt.Sprintf("locale ^= \"%s\"", v))
			break forLoop
		}
	}

	o, err := r.client.ListUsersWithContext(ctx, &cognito.ListUsersInput{
		AttributesToGet: nil,
		Filter:          statement,
		Limit:           aws.Int64(int64(size - 1)),
		PaginationToken: aws.String(token),
		UserPoolId:      aws.String(r.kernel.Config.Cognito.UserPoolID),
	})

	if err != nil {
		return nil, err
	} else if len(o.Users) == 0 {
		return nil, exception.NewCustomError(exception.NotFound, "users")
	}

	users := make([]*domain.User, 0)
	for _, userCg := range o.Users {
		user := mapToUser(userCg)
		users = append(users, user)
	}

	// Add required next pagination token user
	if o.PaginationToken != nil {
		users = append(users, &domain.User{Sub: *o.PaginationToken})
	}

	return users, nil
}

func mapToUser(userCg *cognito.UserType) *domain.User {
	user := &domain.User{
		Sub:        "",
		Email:      "",
		Username:   *userCg.Username,
		Name:       "",
		MiddleName: nil,
		FamilyName: nil,
		Locale:     "",
		Picture:    nil,
		Status:     *userCg.UserStatus,
		CreateTime: userCg.UserCreateDate,
		UpdateTime: userCg.UserLastModifiedDate,
		Enabled:    *userCg.Enabled,
	}

	for _, attr := range userCg.Attributes {
		switch *attr.Name {
		case "sub":
			user.Sub = *attr.Value
			continue
		case "email":
			user.Email = *attr.Value
			continue
		case "name":
			user.Name = *attr.Value
			continue
		case "middle_name":
			user.MiddleName = attr.Value
			continue
		case "family_name":
			user.FamilyName = attr.Value
			continue
		case "locale":
			user.Locale = *attr.Value
			continue
		case "picture":
			user.Picture = attr.Value
			continue
		}
	}

	return user
}