package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/broker"
	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

type handler struct {
	broker *broker.Broker
}

func (h *handler) publishHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading publish request body", err)
		w.Write([]byte("error reading publish request body"))
		return
	}
	defer r.Body.Close()

	var msg message.Message

	if err := json.Unmarshal(body, &msg); err != nil {
		log.Println("error parsing publish message json", err)
		w.Write([]byte("error parsing publish message json"))
		return
	}

	h.broker.Publish(msg)
}

func (h *handler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading subscribe request body", err)
		w.Write([]byte("error reading subscribe request body"))
		return
	}
	defer r.Body.Close()

	var s message.Subscribe

	if err := json.Unmarshal(body, &s); err != nil {
		log.Println("error parsing subscribe json", err)
		w.Write([]byte("error parsing subscribe json"))
		return
	}

	h.broker.Subscribe(s)
}

func (h *handler) unsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading unsubscribe body", err)
		w.Write([]byte("error reading unsubscribe body"))
		return
	}
	defer r.Body.Close()

	var u message.Unsubscribe

	if err := json.Unmarshal(body, &u); err != nil {
		log.Println("error parsing unsubscribe json", err)
		w.Write([]byte("error parsing unsubscribe json"))
		return
	}

	h.broker.Unsubscribe(u)
}

func (h *handler) deleteSubscriberHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading delete subscriber request body", err)
		w.Write([]byte("error reading delete subscriber request body"))
		return
	}
	defer r.Body.Close()

	ID := string(body)

	h.broker.RemoveSubscriber(ID)
}

func (h *handler) addSubscriberHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading add subscriber request body", err)
		w.Write([]byte("error reading add subscriber request body"))
		return
	}
	defer r.Body.Close()

	var s subscriber.Subscriber

	if err := json.Unmarshal(body, &s); err != nil {
		log.Println("error parsing add subscriber json", err)
		w.Write([]byte("error parsing add subscriber json"))
		return
	}

	h.broker.AddSubscriber(&s)

	w.Write([]byte(s.ID))
}

func (h *handler) activateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading activate request body", err)
		w.Write([]byte("error reading activate request body"))
		return
	}
	defer r.Body.Close()

	ID := string(body)

	h.broker.Activate(ID)
}

func (h *handler) deactivateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading activate request body", err)
		w.Write([]byte("error reading activate request body"))
		return
	}
	defer r.Body.Close()

	ID := string(body)

	h.broker.Deactivate(ID)
}
