package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

func main() {
	s := subscriber.New("http://localhost:8081", []string{"football", "volleyball", "waterpolo"}, true)

	j, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	go s.Listen()

	_, err = http.Post("http://localhost:8080/subscribe", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}
