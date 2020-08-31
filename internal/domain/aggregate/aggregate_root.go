package aggregate

import "github.com/alexandria-oss/identity-api/internal/domain/event"

type AggregateRoot interface {
	// Get all the aggregate root's immutable events
	PullDomainEvents() []event.Domain
	// Register a new event
	Record(e event.Domain)
	// Parse binary data to struct
	UnmarshalBinary(data []byte) error
	// Parse current in-memory data to binary
	MarshalBinary() ([]byte, error)
	// Get an aggregate root primitive-only version
	ToPrimitive() interface{}
}
