package broker

import (
	"log"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

func (b *Broker) AddSubscriber(s *subscriber.Subscriber) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Subscribers[s.ID] = s

	log.Printf("subscriber [%s] added\n", s.ID)

	b.subscribe(s, s.SubscribedTopics)
}

func (b *Broker) RemoveSubscriber(ID string) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	for topic, subscribers := range b.TopicTable {
		delete(subscribers, ID)
		log.Printf("[%s] unsubscribed from topic [%s]\n", ID, topic)
	}

	delete(b.Subscribers, ID)

	log.Printf("removed [%s]\n", ID)
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

func (b *Broker) Subscribe(subscribeMessage message.Subscribe) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	subscriber := b.Subscribers[subscribeMessage.ID]
	b.subscribe(subscriber, subscribeMessage.Topics)
}

func (b *Broker) Unsubscribe(unsubscribeMessage message.Unsubscribe) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	for _, t := range unsubscribeMessage.Topics {
		delete(b.TopicTable[t], unsubscribeMessage.ID)
		log.Printf("[%s] unsubscribed from topic [%s]\n", unsubscribeMessage.ID, t)
	}
}

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

func (b *Broker) Activate(ID string) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Subscribers[ID].Active = true
}

func (b *Broker) Deactivate(ID string) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Subscribers[ID].Active = false
}
