package messages

// Message is for relays
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
