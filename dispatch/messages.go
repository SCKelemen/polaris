package dispatch

type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data"`
}

type MessageType string

const (
	Notification MessageType = "notification"
	NewNode      MessageType = "new_node"
)
