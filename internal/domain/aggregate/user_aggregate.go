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

func CreateUserRoot(username value.Username, email value.Email) *UserRoot {
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

func (r UserRoot) PullDomainEvents() []event.Domain {
	return r.events
}

func (r *UserRoot) Record(e event.Domain) {
	r.events = append(r.events, e)
}

func (r UserRoot) ToPrimitive() *UserRootPrimitive {
	return &UserRootPrimitive{User: r.User.ToPrimitive()}
}

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

func (r UserRoot) MarshalBinary() ([]byte, error) {
	return json.Marshal(r.ToPrimitive())
}

func BulkUserToPrimitive(uRoot []*UserRoot) []*UserRootPrimitive {
	users := make([]*UserRootPrimitive, 0)
	for _, u := range uRoot {
		users = append(users, u.ToPrimitive())
	}

	return users
}
