package main

import (
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/broker"
	"github.com/gorilla/mux"
)

const BROKER_SERVER_ADDR = ":8080"

func main() {
	h := handler{broker: broker.New()}

	r := mux.NewRouter()

	r.HandleFunc("/publish", h.publishHandler).Methods("POST")
	r.HandleFunc("/subscribe", h.subscribeHandler).Methods("POST")
	r.HandleFunc("/unsubscribe", h.unsubscribeHandler).Methods("POST")
	r.HandleFunc("/add_subscriber", h.addSubscriberHandler).Methods("POST")
	r.HandleFunc("/delete_subscriber", h.deleteSubscriberHandler).Methods("POST")
	r.HandleFunc("/activate", h.activateHandler).Methods("POST")
	r.HandleFunc("/deactivate", h.deactivateHandler).Methods("POST")

	log.Printf("Starting broker server on [%s]\n", BROKER_SERVER_ADDR)

	if err := http.ListenAndServe(BROKER_SERVER_ADDR, r); err != nil {
		log.Fatal("failed to start broker server on", BROKER_SERVER_ADDR, err)
	}
}
