// Package subscriber contains the type of the subscriber client.
package subscriber

import "github.com/google/uuid"

// Subscriber is the blueprint of the client which registers itself to the broker as a subscriber. TCPAddr is the
// clients TCP address which the broker uses to send back messages when available.
type Subscriber struct {
	ID               string   `json:"ID"`
	TCPAddr          string   `json:"addr"`
	SubscribedTopics []string `json:"subscribed_topics"`
	Active           bool
}

// New constructs and returns a new Subscriber.
func New(TCPAddr string, subscribedTopics []string) *Subscriber {
	ID := uuid.NewString()
	return &Subscriber{ID: ID, TCPAddr: TCPAddr, SubscribedTopics: subscribedTopics, Active: true}
}

// Subscribers is a map of the Subscribers IDs to the Subscribers themselves.
type Subscribers map[string]*Subscriber
