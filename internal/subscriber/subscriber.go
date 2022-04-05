package subscriber

import "github.com/google/uuid"

type Subscriber struct {
	ID               string   `json:"id"`
	HTTPAddr         string   `json:"addr"`
	SubscribedTopics []string `json:"subscribed_topics"`
	Active           bool     `json:"active"`
}

func New(HTTPAddr string, subscribedTopics []string, active bool) *Subscriber {
	ID := uuid.NewString()
	return &Subscriber{ID: ID, HTTPAddr: HTTPAddr, SubscribedTopics: subscribedTopics, Active: active}
}

type Subscribers map[string]*Subscriber
