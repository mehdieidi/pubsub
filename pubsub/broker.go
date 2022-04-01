package pubsub

import (
	"fmt"
	"sync"
)

type Subscribers map[string]*Subscriber

type Broker struct {
	Subscribers Subscribers
	TopicTable  map[string]Subscribers
	Lock        sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		Subscribers: Subscribers{},
		TopicTable:  make(map[string]Subscribers),
	}
}

func (b *Broker) AddSubscriber() *Subscriber {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	s := NewSubscriber()

	b.Subscribers[s.ID] = s

	return s
}

func (b *Broker) Subscribe(s *Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	if b.TopicTable[topic] == nil {
		b.TopicTable[topic] = Subscribers{}
	}

	s.AddTopic(topic)
	b.TopicTable[topic][s.ID] = s

	fmt.Printf("%s Subscribed for topic: %s\n", s.ID, topic)
}

func (b *Broker) Publish(topic string, msg string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	for _, s := range b.TopicTable[topic] {
		m := NewMessage(topic, msg)

		if !s.Active {
			continue
		}

		go func(s *Subscriber, m *Message) {
			s.Signal(m)
		}(s, m)
	}
}

func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	delete(b.TopicTable[topic], s.ID)
	s.RemoveTopic(topic)

	fmt.Printf("%s Unsubscribed from topic: %s\n", s.ID, topic)
}

func (b *Broker) RemoveSubscriber(s *Subscriber) {
	for t := range s.Topics {
		b.Unsubscribe(s, t)
	}

	b.Lock.Lock()
	delete(b.Subscribers, s.ID)
	b.Lock.Unlock()

	s.Destruct()
}

func (b *Broker) Broadcast(msg string, topics []string) {
	for _, t := range topics {
		for _, s := range b.TopicTable[t] {
			m := NewMessage(msg, t)

			go func(s *Subscriber, m *Message) {
				s.Signal(m)
			}(s, m)
		}
	}
}

func (b *Broker) GetSubscribersCount(topic string) int {
	b.Lock.RLock()
	defer b.Lock.RUnlock()
	return len(b.TopicTable[topic])
}
