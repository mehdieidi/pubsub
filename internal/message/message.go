// Package message contains the type which represents the data passed between the broker, publishers, and subscribers.
package message

// Message is the type of the data passing between the broker, publishers, and subscribers.
type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}
