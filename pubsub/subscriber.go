package pubsub

import "sync"

type Subscriber struct {
	ID       string
	Messages chan *Message
	Topics   map[string]bool // Topics subscribed to.
	Active   bool
	Lock     sync.RWMutex
}
