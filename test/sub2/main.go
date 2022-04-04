package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

func main() {
	s := subscriber.New("http://localhost:8082", []string{"football", "volleyball", "handball"}, true)

	j, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	var msg message.Message
	go s.Listen(&msg)

	_, err = http.Post("http://localhost:8080/subscribe", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}
