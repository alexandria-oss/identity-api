package driver

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"sync"
)

var cognitoPool *cognito.CognitoIdentityProvider = nil
var cognitoSingleton = new(sync.Once)

// Generate a new cognito client session using shared credentials at ~/.aws
func NewCognitoSession() *cognito.CognitoIdentityProvider {
	cognitoSingleton.Do(func() {
		cognitoPool = cognito.New(session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable})))
	})

	return cognitoPool
}
