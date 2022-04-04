// Package subscriber contains the type of the subscriber client.
package subscriber

import "github.com/google/uuid"

// Subscriber is the blueprint of the client which registers itself to the broker as a subscriber. HTTPAddr is the
// clients address which the broker uses to send back messages when available.
type Subscriber struct {
	ID               string   `json:"ID"`
	HTTPAddr         string   `json:"addr"`
	SubscribedTopics []string `json:"subscribed_topics"`
	Active           bool     `json:"active"`
}

// New constructs and returns a new Subscriber.
func New(HTTPAddr string, subscribedTopics []string, active bool) *Subscriber {
	ID := uuid.NewString()
	return &Subscriber{ID: ID, HTTPAddr: HTTPAddr, SubscribedTopics: subscribedTopics, Active: active}
}

// Subscribers is a map of the Subscribers IDs to the Subscribers themselves.
type Subscribers map[string]*Subscriber
