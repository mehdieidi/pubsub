package message

type Unsubscribe struct {
	ID     string   `json:"id"`
	Topics []string `json:"topics"`
}
