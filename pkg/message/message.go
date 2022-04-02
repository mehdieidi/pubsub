package message

type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}

func New(topic string, msg string) Message {
	return Message{
		Topic: topic,
		Body:  msg,
	}
}
