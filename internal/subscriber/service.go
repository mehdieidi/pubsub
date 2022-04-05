package subscriber

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/message"
)

func (s *Subscriber) Send(msg message.Message) {
	j, err := json.Marshal(msg)
	if err != nil {
		log.Println("error marshalling message to json", err)
		return
	}

	log.Printf("sending message with topic [%s] to address [%s]\n", msg.Topic, s.HTTPAddr+"/msg")

	_, err = http.Post(s.HTTPAddr+"/msg", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("error sending POST request to address [%s] %v\n", s.HTTPAddr, err)
		return
	}

	log.Printf("message sent to [%s]\n", s.ID)
}
