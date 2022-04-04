package broker

import (
	"log"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

// AddSubscriber registers the given subscriber with the broker. It also calls Subscribe method.
func (b *Broker) AddSubscriber(s *subscriber.Subscriber) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Subscribers[s.ID] = s

	log.Printf("subscriber [%s] added\n", s.ID)

	b.subscribe(s, s.SubscribedTopics)
}

func (b *Broker) subscribe(s *subscriber.Subscriber, topics []string) {
	for _, t := range topics {
		if b.TopicTable[t] == nil {
			b.TopicTable[t] = subscriber.Subscribers{}
		}

		b.TopicTable[t][s.ID] = s

		log.Printf("[%s] subscribed for topic [%s]\n", s.ID, t)
	}
}

// Publish sends the msg to all the subscribers which have been registered with the topic of the msg.
func (b *Broker) Publish(msg message.Message) {
	b.Mutex.RLock()
	defer b.Mutex.RUnlock()

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
