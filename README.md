# pubsub
Distributed Many-to-Many pub sub service.

The broker is a server which handles the publish and subscribe functionalities.

To start the broker server:
```
cd pubsub/cmd
go run .
```

## Subscriber clients

Subscriber clients must have an HTTP address so the broker server can send messages back to this address.

Subscriber clients need to introduce themselves to the broker server and subscribe for topics.

Create new subscriber client and marshall. Example: 
```
s := subscriber.New("http://localhost:8081", []string{"football", "volleyball", "handball"}, true)
j, _ := json.Marshal(s)
```

Clients should start listening before introducing themselves to the server. (In a separate goroutine)
```
go s.Listen()
```

Register with the server and subscribe for the defined topics:
```
http.Post("http://localhost:8080/subscribe", "application/json", bytes.NewBuffer(j))
```

The s.Listen method will handle the incoming messages from the broker server.

## Publishing messages


For publishing messages, publishers only need to send a POST request with the defined json structure to the broker server:
```
// send post request with this json structure to the /publish endpoint:
{
	"topic": "handball",
	"body": "a sample message"
}
```