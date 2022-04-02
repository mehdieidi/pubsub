package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/pkg/broker"
	"github.com/MehdiEidi/pubsub/pkg/message"
)

type Handler struct {
	Broker *broker.Broker
}

func (h *Handler) publishHandler(w http.ResponseWriter, r *http.Request) {
	var msg message.Message

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error reading body"))
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		w.Write([]byte("error parsing json"))
	}

	h.Broker.Publish(msg.Topic, msg.Body)
}

type subscriber struct {
	Topics []string `json:"topics"`
	Addr   string   `json:"addr"`
}

func (h *Handler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var s subscriber

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error reading body"))
	}

	if err := json.Unmarshal(body, &s); err != nil {
		w.Write([]byte("error parsing json"))
	}

	s1 := h.Broker.AddSubscriber()
	s1.Addr = s.Addr

	for _, t := range s.Topics {
		h.Broker.Subscribe(s1, t)
	}

	log.Println("subscriber", s1.ID, "subscribed for", s.Topics)

	w.Write([]byte(s1.ID))
}
