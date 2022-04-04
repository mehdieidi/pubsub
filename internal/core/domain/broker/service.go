package broker

import (
	"log"

	"github.com/MehdiEidi/pubsub/internal/core/domain/message"
	"github.com/MehdiEidi/pubsub/internal/core/domain/subscriber"
	"github.com/google/uuid"
)

func (b *Broker) AddSubscriber(s *subscriber.Subscriber) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	s.ID = uuid.NewString()
	b.Subscribers[s.ID] = s

	log.Printf("subscriber [%s] added\n", s.ID)

	b.Subscribe(s, s.SubscribedTopics)
}

func (b *Broker) Subscribe(s *subscriber.Subscriber, topics []string) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	for _, t := range topics {
		if b.TopicTable[t] == nil {
			b.TopicTable[t] = subscriber.Subscribers{}
		}

		b.TopicTable[t][s.ID] = s

		log.Printf("[%s] subscribed for topic [%s]\n", s.ID, t)
	}
}

func (b *Broker) Publish(msg message.Message) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	for _, s := range b.TopicTable[msg.Topic] {
		if !s.Active {
			continue
		}

		log.Printf("sending message with topic [%s] to subscriber [%s]\n", msg.Topic, s.ID)

		go func(s *subscriber.Subscriber, msg message.Message) {
			s.Send(msg)
		}(s, msg)
	}
}
