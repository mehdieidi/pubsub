package main

import (
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/pkg/broker"
	"github.com/gorilla/mux"
)

const addr = ":8080"

func main() {
	h := Handler{Broker: broker.New()}

	r := mux.NewRouter()

	r.HandleFunc("/publish", h.publishHandler).Methods("POST")
	r.HandleFunc("/subscribe", h.subscribeHandler).Methods("POST")

	log.Println("Server starting with address: ", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
