package subscriber

import (
	"log"
	"sync"

	"github.com/MehdiEidi/pubsub/message"
	"github.com/google/uuid"
)

type Subscriber struct {
	ID               string
	Messages         chan *message.Message
	SubscribedTopics map[string]struct{}
	Active           bool
	Lock             sync.RWMutex
}

func New() *Subscriber {
	ID := uuid.New().String()

	return &Subscriber{
		ID:               ID,
		Messages:         make(chan *message.Message),
		SubscribedTopics: make(map[string]struct{}),
		Active:           true,
	}
}

func (s *Subscriber) Subscribe(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.SubscribedTopics[topic] = struct{}{}
}

func (s *Subscriber) Unsubscribe(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.SubscribedTopics, topic)
}

func (s *Subscriber) Topics() []string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	var topics []string

	for t := range s.SubscribedTopics {
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
					log.Printf("[Subscriber %s] received message %s. [topic %s]\n", s.ID, msg.Body, msg.Topic)
				}
			}
		}
	}()
}

func (s *Subscriber) ReceiveMessage(msg *message.Message) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	s.Messages <- msg
}

func (s *Subscriber) Delete() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Active = false
	close(s.Messages)
}
