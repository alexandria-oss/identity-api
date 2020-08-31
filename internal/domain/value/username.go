package value

import "github.com/alexandria-oss/common-go/exception"

type Username struct {
	Value string
}

func NewUsernameFromString(username string) (*Username, error) {
	value := &Username{Value: username}
	if err := value.IsValid(); err != nil {
		return nil, err
	}

	return value, nil
}

func (u *Username) Rename(username string) {
	u.Value = username
}

func (u Username) IsValid() error {
	if len(u.Value) == 0 && len(u.Value) > 256 {
		return exception.NewFieldRange("username", "1", "256")
	}

	return nil
}
