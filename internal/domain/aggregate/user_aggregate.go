package aggregate

import (
	"github.com/alexandria-oss/identity-api/internal/domain/entity"
	"github.com/alexandria-oss/identity-api/internal/domain/event"
)

type UserRoot struct {
	Root         entity.User
	domainEvents []event.Domain
}

func (r *UserRoot) PullDomainEvents() []event.Domain {
	return r.domainEvents
}

func (r *UserRoot) Record(e event.Domain) {
	r.domainEvents = append(r.domainEvents, e)
}
