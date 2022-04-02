package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/pkg/broker"
	"github.com/MehdiEidi/pubsub/pkg/message"
)

type subscriberClient struct {
	Topics []string `json:"topics"`
	Addr   string   `json:"addr"`
}
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

	log.Println("publishing", msg.Topic, msg.Body)

	h.Broker.Publish(msg.Topic, msg.Body)
}

func (h *Handler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var client subscriberClient

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error reading body"))
	}

	if err := json.Unmarshal(body, &client); err != nil {
		w.Write([]byte("error parsing json"))
	}

	s := h.Broker.AddSubscriber()
	s.Addr = client.Addr

	for _, t := range client.Topics {
		h.Broker.Subscribe(s, t)
	}

	w.Write([]byte(s.ID))
}
