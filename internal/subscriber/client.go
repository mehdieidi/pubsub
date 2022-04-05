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

type handler struct {
	msg *message.Message
}

func (s *Subscriber) Listen(msg *message.Message) {
	h := handler{msg: msg}

	r := mux.NewRouter()

	r.HandleFunc("/msg", h.messageReceiveHandler).Methods("POST")

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

func (h *handler) messageReceiveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading message body", err)
		w.Write([]byte("error reading message body"))
		return
	}
	defer r.Body.Close()

	var msg message.Message

	if err := json.Unmarshal(body, &msg); err != nil {
		log.Println("error parsing message json", err)
		w.Write([]byte("error parsing message json"))
		return
	}

	log.Printf("received [%s] %s\n", msg.Topic, msg.Body)

	*h.msg = msg
}
