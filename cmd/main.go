package main

import (
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/core/domain/broker"
	"github.com/gorilla/mux"
)

const BROKER_ADDR = ":8080"

func main() {
	h := handler{broker: broker.New()}

	r := mux.NewRouter()

	r.HandleFunc("/publish", h.publishHandler).Methods("POST")
	r.HandleFunc("subscriber", h.subscribeHandler).Methods("POST")

	log.Printf("Starting broker server on %s", BROKER_ADDR)

	if err := http.ListenAndServe(BROKER_ADDR, r); err != nil {
		log.Fatal("failed to start broker server on", BROKER_ADDR, err)
	}
}
