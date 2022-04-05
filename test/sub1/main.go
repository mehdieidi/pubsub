package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MehdiEidi/pubsub/internal/message"
	"github.com/MehdiEidi/pubsub/internal/subscriber"
)

func main() {
	s := subscriber.New("http://localhost:8081", []string{"football", "volleyball", "waterpolo"}, true)

	j, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	var msg message.Message
	go s.Listen(&msg)

	resp, err := http.Post("http://localhost:8080/add_subscriber", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	ID := string(responseBody)

	fmt.Println(ID)

	time.Sleep(5 * time.Second)

	subscribeMsg := message.Subscribe{ID: ID, Topics: []string{"handball"}}

	j, err = json.Marshal(subscribeMsg)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Post("http://localhost:8080/subscribe", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}
