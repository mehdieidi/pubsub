package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/pkg/message"
	"github.com/gorilla/mux"
)

func Listen(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/", msgHandler).Methods("POST")

	fmt.Println("listening for messages on", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got something")

	var msg message.Message

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error reading body"))
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		w.Write([]byte("error parsing json"))
	}

	fmt.Println("Got message", msg.Topic, msg.Body)
}
