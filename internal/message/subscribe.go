package message

type Subscribe struct {
	ID     string   `json:"id"`
	Topics []string `json:"topics"`
}
