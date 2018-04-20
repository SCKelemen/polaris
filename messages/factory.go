package messages

import (
	"encoding/json"
	"fmt"
)

// MessageFactory is a way to manage the creation of messages
type MessageFactory struct {
}

// New creates a new message factory
func New() MessageFactory {
	return MessageFactory{}
}

// func (mf MessageFactory) CreateMessage(messageType string, data string) Message {
// 	switch messageType {
// 	}
// }

// CreateJoinRequest creates the request for when a node wants to join the graph
func (mf MessageFactory) CreateJoinRequest(name string, address string) Message {
	msg := Message{Type: "new_node"}
	data := NewNodeData{Name: name, Address: address}
	json, err := json.Marshal(data)
	if err == nil {
		msg.Data = string(json)
	} else {
		fmt.Println("failed to serialize data")
	}
	return msg
}

type NewNodeData struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}
