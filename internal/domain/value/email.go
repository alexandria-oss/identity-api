package value

import (
	"github.com/alexandria-oss/common-go/exception"
	"strings"
)

type Email struct {
	Value string
}

func NewEmailFromString(email string) (*Email, error) {
	value := &Email{Value: email}
	if err := value.IsValid(); err != nil {
		return nil, err
	}

	return value, nil
}

func (e *Email) Change(email string) error {
	snapshot := e.Value
	e.Value = email

	if err := e.IsValid(); err != nil {
		e.Value = snapshot
		return err
	}

	return nil
}

func (e Email) IsValid() error {
	switch {
	case len(e.Value) == 0 && len(e.Value) > 512:
		return exception.NewFieldRange("email", "1", "512")
	case !strings.Contains(e.Value, "@"):
		return exception.NewFieldFormat("email", "valid email")
	default:
		return nil
	}
}
