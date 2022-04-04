package subscriber

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/gorilla/mux"
)

// Listen starts an http server on the subscribers address and listens for incoming messages.
func (s *Subscriber) Listen() {
	r := mux.NewRouter()

	r.HandleFunc("/msg", messageReceiveHandler).Methods("POST")

	u, err := url.Parse(s.HTTPAddr)
	if err != nil {
		log.Printf("error parsing HTTP address [%s] of subscriber [%s]\n", s.HTTPAddr, s.ID)
		return
	}

	u.Scheme = ""
	TCPAddr := u.String()[2:]

	log.Printf("Subscriber [%s] starting to listen to [%s]\n", s.ID, TCPAddr)

	if err := http.ListenAndServe(TCPAddr, r); err != nil {
		log.Printf("error starting server for subscriber [%s] on address [%s] %v\n", s.ID, TCPAddr, err)
	}
}

// messageReceiveHandler handles the incoming messages to the subscriber.
func messageReceiveHandler(w http.ResponseWriter, r *http.Request) {
	var msg message.Message

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading message body", err)
		w.Write([]byte("error reading message body"))
		return
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		log.Println("error parsing json body", err)
		w.Write([]byte("error parsing json body"))
		return
	}

	log.Printf("received [%s] %s\n", msg.Topic, msg.Body)
}
