package driver

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// Generate a new cognito client session using shared credentials at ~/.aws
func NewCognitoSession() *cognito.CognitoIdentityProvider {
	return cognito.New(session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable})))
}
