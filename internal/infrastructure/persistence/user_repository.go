package persistence

import (
	"context"
	"github.com/alexandria-oss/identity-api/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"sync"
)

type UserCommandAWSRepository struct {
	client *cognito.CognitoIdentityProvider
	kernel domain.KernelStore
	mu     *sync.RWMutex
}

func NewUserCommandAWSRepository(c *cognito.CognitoIdentityProvider, k domain.KernelStore) *UserCommandAWSRepository {
	return &UserCommandAWSRepository{
		client: c,
		kernel: k,
		mu:     new(sync.RWMutex),
	}
}

func (r *UserCommandAWSRepository) Remove(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminDisableUserWithContext(ctx, &cognito.AdminDisableUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})

	return err
}

func (r *UserCommandAWSRepository) Restore(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminEnableUserWithContext(ctx, &cognito.AdminEnableUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})

	return err
}

func (r *UserCommandAWSRepository) HardRemove(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.client.AdminDeleteUserWithContext(ctx, &cognito.AdminDeleteUserInput{
		UserPoolId: aws.String(r.kernel.Config.Cognito.UserPoolID),
		Username:   aws.String(id),
	})

	return err
}
