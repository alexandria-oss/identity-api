package event

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Domain struct {
	AggregateID string    `json:"aggregate_id"`
	ID          uuid.UUID `json:"id"`
	Service     string    `json:"service"`
	Action      string    `json:"action"`
	Body        []byte    `json:"body"`
	PublishTime time.Time `json:"publish_time"`
}

func NewDomain(aggregateID, service, action string, body []byte) *Domain {
	return &Domain{
		AggregateID: aggregateID,
		ID:          uuid.New(),
		Service:     strings.ToLower(service),
		Action:      strings.ToLower(action),
		Body:        body,
		PublishTime: time.Now(),
	}
}

func (d Domain) GetName() string {
	return fmt.Sprintf("%s.%s", d.Service, d.Action)
}
