package pubsub

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Subscriber struct {
	ID       string
	Messages chan *Message
	Topics   map[string]bool // Topics subscribed to.
	Active   bool
	Lock     sync.RWMutex
}

func NewSubscriber() *Subscriber {
	s := rand.NewSource(int64(time.Now().UnixNano()))
	r := rand.New(s)
	ID := strconv.Itoa(r.Intn(200))

	return &Subscriber{
		ID:       ID,
		Messages: make(chan *Message),
		Topics:   make(map[string]bool),
		Active:   true,
	}
}

func (s *Subscriber) AddTopic(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Topics[topic] = true
}

func (s *Subscriber) RemoveTopic(topic string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Topics, topic)
}

func (s *Subscriber) GetTopics() []string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	var topics []string

	for t := range s.Topics {
		topics = append(topics, t)
	}

	return topics
}

func (s *Subscriber) Signal(msg *Message) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	if s.Active {
		s.Messages <- msg
	}
}

func (s *Subscriber) Destruct() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Active = false
	close(s.Messages)
}

func (s *Subscriber) Listen() {
	for {
		if msg, ok := <-s.Messages; ok {
			fmt.Printf("Subscriber %s, received %s with topic %s\n", s.ID, msg.Body, msg.Topic)
		}
	}
}
