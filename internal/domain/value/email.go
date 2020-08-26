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

func (e *Email) Change(email string) {
	e.Value = email
}

func (e Email) IsValid() error {
	switch {
	case len(e.Value) == 0 && len(e.Value) > 512:
		return exception.NewCustomError(exception.FieldRange, "email", "1", "512")
	case !strings.Contains(e.Value, "@"):
		return exception.NewCustomError(exception.FieldFormat, "email", "valid email")
	default:
		return nil
	}
}
