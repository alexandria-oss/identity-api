package event

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Domain struct {
	ID            uuid.UUID `json:"id"`
	AggregateID   string    `json:"aggregate_id"`
	Service       string    `json:"service"`
	AggregateName string    `json:"aggregate_name"`
	Action        string    `json:"action"`
	Body          []byte    `json:"body"`
	Snapshot      []byte    `json:"snapshot"`
	PublishTime   time.Time `json:"publish_time"`
}

func NewDomain(aggregateID, service, aggregateName, action string, body, snapshot []byte) Domain {
	return Domain{
		ID:            uuid.New(),
		AggregateID:   aggregateID,
		Service:       strings.ToLower(service),
		AggregateName: strings.ToLower(aggregateName),
		Action:        strings.ToLower(action),
		Body:          body,
		Snapshot:      snapshot,
		PublishTime:   time.Now(),
	}
}

func (d Domain) GetName() string {
	return fmt.Sprintf("%s.%s.%s", d.Service, d.AggregateName, d.Action)
}
