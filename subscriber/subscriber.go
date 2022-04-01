package subscriber

import (
	"log"
	"sync"

	"github.com/MehdiEidi/gods/set"
	"github.com/MehdiEidi/pubsub/message"
	"github.com/google/uuid"
)

type Subscriber struct {
	ID               string
	Messages         chan message.Message
	SubscribedTopics *set.Set
	Active           bool
	Lock             sync.RWMutex
}

func New() *Subscriber {
	ID := uuid.New().String()

	return &Subscriber{
		ID:               ID,
		Messages:         make(chan message.Message),
		SubscribedTopics: set.New(),
		Active:           true,
	}
}

func (s *Subscriber) Subscribe(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.SubscribedTopics.Add(topic)
}

func (s *Subscriber) Unsubscribe(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.SubscribedTopics.Delete(topic)
}

func (s *Subscriber) Topics() []string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	var topics []string

	for t := range s.SubscribedTopics.Data {
		topics = append(topics, t)
	}

	return topics
}

func (s *Subscriber) Listen() {
	go func() {
		for {
			select {
			case msg, ok := <-s.Messages:
				if ok {
					log.Printf("[%s] received message %s. [topic %s]\n", s.ID, msg.Body, msg.Topic)
				}
			}
		}
	}()
}

func (s *Subscriber) Send(msg message.Message) {
	s.Messages <- msg
}

func (s *Subscriber) Delete() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Active = false
	close(s.Messages)
}
