package broker

import (
	"log"
	"sync"

	"github.com/MehdiEidi/pubsub/message"
	"github.com/MehdiEidi/pubsub/subscriber"
)

type subscribers map[string]*subscriber.Subscriber

type Broker struct {
	Subscribers subscribers
	TopicTable  map[string]subscribers
	Lock        sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		Subscribers: subscribers{},
		TopicTable:  make(map[string]subscribers),
	}
}

func (b *Broker) AddSubscriber() *subscriber.Subscriber {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	s := subscriber.New()
	b.Subscribers[s.ID] = s

	return s
}

func (b *Broker) RemoveSubscriber(s *subscriber.Subscriber) {
	for t := range s.SubscribedTopics {
		b.Unsubscribe(s, t)
	}

	b.Lock.Lock()
	delete(b.Subscribers, s.ID)
	b.Lock.Unlock()

	s.Delete()
}

func (b *Broker) SubscribersCount(topic string) int {
	b.Lock.RLock()
	defer b.Lock.RUnlock()
	return len(b.TopicTable[topic])
}

func (b *Broker) Subscribe(s *subscriber.Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	if b.TopicTable[topic] == nil {
		b.TopicTable[topic] = subscribers{}
	}

	s.Subscribe(topic)
	b.TopicTable[topic][s.ID] = s

	log.Printf("[%s] subscribed for topic [%s]\n", s.ID, topic)
}

func (b *Broker) Unsubscribe(s *subscriber.Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	delete(b.TopicTable[topic], s.ID)
	s.Unsubscribe(topic)

	log.Printf("[%s] unsubscribed from topic [%s]\n", s.ID, topic)
}

func (b *Broker) Publish(topic string, msg string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	for _, s := range b.TopicTable[topic] {
		m := message.New(topic, msg)

		if !s.Active {
			continue
		}

		go func(s *subscriber.Subscriber, m *message.Message) {
			s.ReceiveMessage(m)
		}(s, m)
	}
}

func (b *Broker) Broadcast(msg string, topics []string) {
	for _, t := range topics {
		for _, s := range b.TopicTable[t] {
			m := message.New(msg, t)

			go func(s *subscriber.Subscriber, m *message.Message) {
				s.ReceiveMessage(m)
			}(s, m)
		}
	}
}
