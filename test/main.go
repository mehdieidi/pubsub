package main

import (
	"fmt"

	"github.com/MehdiEidi/pubsub/pkg/broker"
)

func main() {
	topics := []string{"football", "volleyball", "swimming", "handball", "waterpolo"}

	b := broker.New()

	s1 := b.AddSubscriber()
	s2 := b.AddSubscriber()
	s3 := b.AddSubscriber()

	b.Subscribe(s1, topics[1])
	b.Subscribe(s1, topics[3])
	b.Subscribe(s1, topics[4])

	b.Subscribe(s2, topics[1])
	b.Subscribe(s2, topics[2])

	b.Subscribe(s3, topics[1])
	b.Subscribe(s3, topics[0])

	s1.Listen()
	s2.Listen()
	s3.Listen()

	b.Publish(topics[0], "f1")

	b.Publish(topics[1], "v1")
	b.Publish(topics[1], "v2")

	b.Publish(topics[0], "f2")

	b.Publish(topics[2], "s1")

	b.Publish(topics[3], "h1")

	b.Publish(topics[4], "w1")

	b.Unsubscribe(s2, topics[1])

	go b.Publish(topics[1], "v3")
	go b.Publish(topics[1], "v4")

	b.Multicast("multi", []string{topics[0], topics[4]})

	b.Broadcast("broad")

	fmt.Scanln()
}
