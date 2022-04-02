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

	//test code {
	// topics := []string{"football", "volleyball", "swimming", "handball", "waterpolo"}

	// s1 := h.Broker.AddSubscriber()
	// s2 := h.Broker.AddSubscriber()
	// s3 := h.Broker.AddSubscriber()

	// h.Broker.Subscribe(s1, topics[1])
	// h.Broker.Subscribe(s1, topics[3])
	// h.Broker.Subscribe(s1, topics[4])

	// h.Broker.Subscribe(s2, topics[1])
	// h.Broker.Subscribe(s2, topics[2])

	// h.Broker.Subscribe(s3, topics[1])
	// h.Broker.Subscribe(s3, topics[0])

	// s1.Listen()
	// s2.Listen()
	// s3.Listen()
	// }

	r := mux.NewRouter()

	r.HandleFunc("/publish", h.publishHandler).Methods("POST")
	r.HandleFunc("/subscribe", h.subscribeHandler).Methods("POST")

	log.Println("Server starting with address: ", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
