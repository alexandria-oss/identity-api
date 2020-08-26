package persistence

import (
	"context"
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"strings"
	"sync"
)

type UserAWSRepository struct {
	client *cognito.CognitoIdentityProvider
	kernel domain.KernelStore
	mu     *sync.RWMutex
}

func NewUserAWSRepository(c *cognito.CognitoIdentityProvider, k domain.KernelStore) *UserAWSRepository {
	return &UserAWSRepository{
		client: c,
		kernel: k,
		mu:     new(sync.RWMutex),
	}
}

func (r *UserAWSRepository) Remove(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminDisableUserWithContext(ctx, &cognito.AdminDisableUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "NotFound"):
			return exception.NewCustomError(exception.NotFound, "user")
		default:
			return err
		}
	}

	return nil
}

func (r *UserAWSRepository) Restore(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminEnableUserWithContext(ctx, &cognito.AdminEnableUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "NotFound"):
			return exception.NewCustomError(exception.NotFound, "user")
		default:
			return err
		}
	}

	return nil
}

func (r *UserAWSRepository) HardRemove(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminDeleteUserWithContext(ctx, &cognito.AdminDeleteUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "NotFound"):
			return exception.NewCustomError(exception.NotFound, "user")
		default:
			return err
		}
	}

	return nil
}

func (r *UserAWSRepository) FetchOne(ctx context.Context, byUsername bool, key string) (*aggregate.UserRoot, error) {
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

	return r.mapToUser(o.Users[0])
}

func (r *UserAWSRepository) Fetch(ctx context.Context, criteria domain.Criteria) ([]*aggregate.UserRoot, domain.PaginationToken, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tokenStr *string = nil
	if criteria.Token != "" {
		tokenStr = aws.String(string(criteria.Token))
	}

	var statement *string = nil
forLoop:
	for k, v := range criteria.FilterBy {
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
		case k == "disabled" && v != "":
			statement = aws.String(fmt.Sprintf("enabled = false"))
			break forLoop
		}
	}

	o, err := r.client.ListUsersWithContext(ctx, &cognito.ListUsersInput{
		AttributesToGet: nil,
		Filter:          statement,
		Limit:           aws.Int64(int64(criteria.Limit)),
		PaginationToken: tokenStr,
		UserPoolId:      aws.String(r.kernel.Config.Cognito.UserPoolID),
	})

	if err != nil {
		return nil, "", err
	} else if len(o.Users) == 0 {
		return nil, "", exception.NewCustomError(exception.NotFound, "users")
	}

	users := make([]*aggregate.UserRoot, 0)
	for _, userCg := range o.Users {
		user, err := r.mapToUser(userCg)
		if err != nil {
			return nil, "", err
		}
		users = append(users, user)
	}

	nextToken := domain.PaginationToken("")
	if o.PaginationToken != nil {
		nextToken = domain.PaginationToken(*o.PaginationToken)
	}

	return users, nextToken, nil
}

// Adapter
func (r UserAWSRepository) mapToUser(userCg *cognito.UserType) (*aggregate.UserRoot, error) {
	userPrim := &entity.UserPrimitive{
		ID:         "",
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
			userPrim.ID = *attr.Value
			continue
		case "email":
			userPrim.Email = *attr.Value
			continue
		case "name":
			userPrim.Name = *attr.Value
			continue
		case "middle_name":
			userPrim.MiddleName = attr.Value
			continue
		case "family_name":
			userPrim.FamilyName = attr.Value
			continue
		case "locale":
			userPrim.Locale = *attr.Value
			continue
		case "picture":
			userPrim.Picture = attr.Value
			continue
		}
	}

	user, err := userPrim.ToEntity()
	if err != nil {
		return nil, err
	}

	return &aggregate.UserRoot{User: user}, nil
}
