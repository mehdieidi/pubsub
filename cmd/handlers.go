package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/core/domain/broker"
	"github.com/MehdiEidi/pubsub/internal/core/domain/message"
	"github.com/MehdiEidi/pubsub/internal/core/domain/subscriber"
)

type handler struct {
	broker *broker.Broker
}

func (h *handler) publishHandler(w http.ResponseWriter, r *http.Request) {
	var msg message.Message

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading publish request body", err)
		w.Write([]byte("error reading publish request body"))
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		log.Println("error parsing publish message json body", err)
		w.Write([]byte("error parsing publish message json body"))
	}

	h.broker.Publish(msg)
}

func (h *handler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var s subscriber.Subscriber

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading subscriber request body", err)
		w.Write([]byte("error reading subscriber request body"))
	}

	if err := json.Unmarshal(body, &s); err != nil {
		log.Println("error parsing subscriber json body", err)
		w.Write([]byte("error parsing subscriber json body"))
	}

	h.broker.AddSubscriber(&s)

	w.Write([]byte(s.ID))
}
