// Package broker contains the blueprint of the broker server which is responsible for getting messages from the
// publishers and sending those messages to the subscribers.
package broker

import (
	"sync"

	"github.com/MehdiEidi/pubsub/internal/core/domain/subscriber"
)

// Broker is the broker server which is responsible for getting published messages and sending those messages to the
// subscribers.
type Broker struct {
	Subscribers subscriber.Subscribers            // All subscribers registered to the broker.
	TopicTable  map[string]subscriber.Subscribers // Mapping topics to their subscribers.
	Mutex       sync.RWMutex
}

// New constructs and returns a new Broker.
func New() *Broker {
	return &Broker{
		Subscribers: subscriber.Subscribers{},
		TopicTable:  make(map[string]subscriber.Subscribers),
	}
}
