package aggregate

import (
	"encoding/json"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	"github.com/alexandria-oss/identity-api/internal/domain/event"
	"github.com/alexandria-oss/identity-api/internal/domain/value"
)

type UserRoot struct {
	User   *entity.User `json:"user"`
	events []event.Domain
}

type UserRootPrimitive struct {
	User *entity.UserPrimitive `json:"user"`
}

func NewUserRoot(username value.Username, email value.Email) *UserRoot {
	// Gen default value objects
	userID := &value.UserID{Value: ""}

	state := new(value.State)
	state.Activate()

	user := &UserRoot{
		User: &entity.User{
			ID:         userID,
			Username:   &username,
			Email:      &email,
			Name:       "",
			MiddleName: nil,
			FamilyName: nil,
			Locale:     "",
			Picture:    nil,
			Status:     "",
			CreateTime: nil,
			UpdateTime: nil,
			Enabled:    state,
		},
		events: nil,
	}

	user.Record(event.NewDomain(
		user.User.Username.Value,
		"identity",
		"user",
		"create",
		nil,
		nil,
	))
	return user
}

// PullDomainEvents Get all the aggregate root's immutable events
func (r UserRoot) PullDomainEvents() []event.Domain {
	return r.events
}

// Record Register a new event
func (r *UserRoot) Record(e event.Domain) {
	r.events = append(r.events, e)
}

// ToPrimitive Get an aggregate root primitive-only version
func (r UserRoot) ToPrimitive() *UserRootPrimitive {
	return &UserRootPrimitive{User: r.User.ToPrimitive()}
}

// UnmarshalBinary Parse binary data to struct
func (r *UserRoot) UnmarshalBinary(data []byte) error {
	prim := new(UserRootPrimitive)
	if err := json.Unmarshal(data, prim); err != nil {
		return err
	}

	u, err := prim.User.ToEntity()
	if err != nil {
		return err
	}

	r.User = u
	return nil
}

// MarshalBinary Parse current in-memory data to binary
func (r UserRoot) MarshalBinary() ([]byte, error) {
	return json.Marshal(r.ToPrimitive())
}
