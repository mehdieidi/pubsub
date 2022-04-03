package domain

type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}
