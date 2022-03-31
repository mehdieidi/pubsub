package pubsub

import "sync"

type Subscribers map[string]Subscriber

type Broker struct {
	Subscribers Subscribers
	Topics      map[string]Subscribers
	Lock        sync.RWMutex
}

func (b *Broker) Subscribe(s *Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	if b.Topics[topic] == nil {
		b.Topics[topic] = Subscribers{}
	}

	s.AddTopic(topic)
	b.Topics[topic][s.ID] = *s
}

func (b *Broker) Publish(topic string, msg string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	for _, s := range b.Topics[topic] {
		m := NewMessage(msg, topic)

		if !s.Active {
			continue
		}

		go func(s *Subscriber) {
			s.Signal(m)
		}(&s)
	}
}

func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	delete(b.Topics[topic], s.ID)
	s.RemoveTopic(topic)
}
