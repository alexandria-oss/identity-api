package infrastructure

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/user"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"sync"
)

type UserQueryAWSRepository struct {
	client *cognito.CognitoIdentityProvider
	kernel domain.KernelStore
	mu     *sync.RWMutex
}

// Factory Method
func NewUserQueryAWSRepository(c *cognito.CognitoIdentityProvider, k domain.KernelStore) *UserQueryAWSRepository {
	return &UserQueryAWSRepository{
		client: c,
		kernel: k,
		mu:     new(sync.RWMutex),
	}
}

func (r *UserQueryAWSRepository) FetchOne(ctx context.Context, byUsername bool, key string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	statement := aws.String(fmt.Sprintf("sub = \"%s\"", key))
	if byUsername {
		statement = aws.String(fmt.Sprintf("username = \"%s\"", key))
	}

	o, err := r.client.ListUsersWithContext(ctx, &cognito.ListUsersInput{
		AttributesToGet: nil,
		Filter:          statement,
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

func (r *UserQueryAWSRepository) Fetch(ctx context.Context, token string, size int, filterMap domain.FilterMap) ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tokenStr *string = nil
	if token != "" {
		tokenStr = aws.String(token)
	}

	var statement *string = nil
forLoop:
	for k, v := range filterMap {
		switch {
		case k == "name" && v != "":
			statement = aws.String(fmt.Sprintf("name ^= \"%s\"", v))
			break forLoop
		case k == "email" && v != "":
			statement = aws.String(fmt.Sprintf("email ^= \"%s\"", v))
			break forLoop
		case k == "middle_name" && v != "":
			statement = aws.String(fmt.Sprintf("middle_name ^= \"%s\"", v))
			break forLoop
		case k == "family_name" && v != "":
			statement = aws.String(fmt.Sprintf("family_name ^= \"%s\"", v))
			break forLoop
		case k == "locale" && v != "":
			statement = aws.String(fmt.Sprintf("locale ^= \"%s\"", v))
			break forLoop
		}
	}

	o, err := r.client.ListUsersWithContext(ctx, &cognito.ListUsersInput{
		AttributesToGet: nil,
		Filter:          statement,
		Limit:           aws.Int64(int64(size - 1)),
		PaginationToken: tokenStr,
		UserPoolId:      aws.String(r.kernel.Config.Cognito.UserPoolID),
	})

	if err != nil {
		return nil, err
	} else if len(o.Users) == 0 {
		return nil, exception.NewCustomError(exception.NotFound, "users")
	}

	users := make([]*user.User, 0)
	for _, userCg := range o.Users {
		user := mapToUser(userCg)
		users = append(users, user)
	}

	// Add required next pagination token user
	if o.PaginationToken != nil && o.PaginationToken != aws.String("") {
		users = append(users, &user.User{Sub: *o.PaginationToken})
	}

	return users, nil
}

func mapToUser(userCg *cognito.UserType) *user.User {
	user := &user.User{
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
