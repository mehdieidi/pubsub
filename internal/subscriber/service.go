package subscriber

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/gorilla/mux"
)

// Listen starts an http server on the subscribers address and listens for incoming messages.
func (s *Subscriber) Listen() {
	r := mux.NewRouter()

	r.HandleFunc("/msg", messageReceiveHandler).Methods("POST")

	log.Printf("Subscriber [%s] listening to [%s]\n", s.ID, s.TCPAddr)

	if err := http.ListenAndServe(s.TCPAddr, r); err != nil {
		log.Printf("error starting listener server for subscriber [%s] on address [%s]\n", s.ID, s.TCPAddr)
	}
}

// messageReceiveHandler handles the incoming messages.
func messageReceiveHandler(w http.ResponseWriter, r *http.Request) {
	var msg message.Message

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading message body", err)
		w.Write([]byte("error reading message body"))
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		log.Println("error parsing json body", err)
		w.Write([]byte("error parsing json body"))
	}

	log.Printf("[%s] %s\n", msg.Topic, msg.Body)
}

func (s *Subscriber) Send(msg message.Message) {
	j, err := json.Marshal(msg)
	if err != nil {
		log.Println("error marshalling message to json", err)
		return
	}

	log.Printf("sending message with topic [%s] to address [%s]\n", msg.Topic, s.TCPAddr)

	_, err = http.Post(s.TCPAddr, "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("error sending POST request to address [%s]\n %v", s.TCPAddr, err)
		return
	}

	log.Printf("message sent to [%s]\n", s.ID)
}
