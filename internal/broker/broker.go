package broker

import (
	"sync"

	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

type Broker struct {
	Subscribers subscriber.Subscribers            // All subscribers registered to the broker.
	TopicTable  map[string]subscriber.Subscribers // Mapping topics to their subscribers.
	Mutex       sync.RWMutex
}

func New() *Broker {
	return &Broker{
		Subscribers: subscriber.Subscribers{},
		TopicTable:  make(map[string]subscriber.Subscribers),
	}
}
