package domain

import (
	"sync"
)

type subscribers map[string]*Subscriber

type Broker struct {
	Subscribers subscribers            // All subscribers.
	TopicTable  map[string]subscribers // Map topics to the subscribers of that topic.
	Lock        sync.RWMutex
}
