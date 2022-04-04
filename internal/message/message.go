// Package message contains the type which represents the data passed between broker, publishers, and subscribers.
package message

// Message is the type of the data passing between broker, publishers, and subscribers.
type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}

// New constructs and returns a Message out of the given topic and body.
func New(topic string, body string) Message {
	return Message{Topic: topic, Body: body}
}
