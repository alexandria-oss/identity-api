package mapper

import (
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// AWS Cognito API Adapters - Adapters

// UserAWSToDomain Convert AWS User to Aggregate root
func UserAWSToDomain(userCg *cognito.UserType) (*aggregate.UserRoot, error) {
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
