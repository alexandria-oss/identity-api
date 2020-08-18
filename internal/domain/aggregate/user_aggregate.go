package aggregate

import (
	"encoding/json"
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	"github.com/alexandria-oss/identity-api/internal/domain/event"
)

type UserRoot struct {
	Root         *entity.User `json:"root"`
	domainEvents []event.Domain
}

func (r *UserRoot) PullDomainEvents() []event.Domain {
	return r.domainEvents
}

func (r *UserRoot) Record(e event.Domain) {
	r.domainEvents = append(r.domainEvents, e)
}

func (r *UserRoot) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *UserRoot) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *UserRoot) GetRoot() interface{} {
	return r.Root
}
