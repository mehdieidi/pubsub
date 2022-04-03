package domain

import (
	"sync"

	"github.com/MehdiEidi/gods/set"
)

type Subscriber struct {
	ID               string
	Addr             string
	Messages         chan Message
	SubscribedTopics *set.Set
	Active           bool
	Lock             sync.RWMutex
}
