package broker

import (
	"log"
	"sync"

	"github.com/MehdiEidi/pubsub/message"
	"github.com/MehdiEidi/pubsub/subscriber"
)

const BroadcastTopic = "broadcast"

// subscribers maps subscriber IDs to corresponding Subscriber.
type subscribers map[string]*subscriber.Subscriber

type Broker struct {
	Subscribers subscribers            // All subscribers.
	TopicTable  map[string]subscribers // Map topics to the subscribers of that topic.
	Lock        sync.RWMutex
}

func New() *Broker {
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
	for t := range s.SubscribedTopics.Data {
		b.Unsubscribe(s, t)
	}

	b.Lock.Lock()
	delete(b.Subscribers, s.ID)
	b.Lock.Unlock()

	s.Delete()
}

// TODO add remove subscriber b.add(s)

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

// TODO change params to get Message instead of topic msg

func (b *Broker) Publish(topic string, msg string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	m := message.New(topic, msg)

	for _, s := range b.TopicTable[topic] {
		if !s.Active {
			continue
		}

		go func(s *subscriber.Subscriber, m message.Message) {
			s.Send(m)
		}(s, m)
	}
}

func (b *Broker) Multicast(msg string, topics []string) {
	for _, t := range topics {
		m := message.New(t, msg)

		for _, s := range b.TopicTable[t] {
			go func(s *subscriber.Subscriber, m message.Message) {
				s.Send(m)
			}(s, m)
		}
	}
}

func (b *Broker) Broadcast(msg string) {
	m := message.New(BroadcastTopic, msg)

	for _, s := range b.Subscribers {
		go func(s *subscriber.Subscriber, m message.Message) {
			s.Send(m)
		}(s, m)
	}
}
