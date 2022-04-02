package main

import "encoding/json"

func main() {
	s := []byte(`{"topics":["football", "volleybal"], "addr": ":8081"}`)

	j, err := json.Marshal(s)
}
