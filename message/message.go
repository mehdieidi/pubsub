package message

type Message struct {
	Topic string
	Body  string
}

func New(topic string, msg string) *Message {
	return &Message{
		Topic: topic,
		Body:  msg,
	}
}
