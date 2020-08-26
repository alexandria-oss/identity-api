package aggregate

import "github.com/alexandria-oss/identity-api/internal/domain/event"

type AggregateRoot interface {
	PullDomainEvents() []event.Domain
	Record(e event.Domain)
	UnmarshalBinary(data []byte) error
	MarshalBinary() ([]byte, error)
	GetRoot() interface{}
}
