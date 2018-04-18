package messages

import (
	"github.com/SCKelemen/polaris/dispatch"
)

const (
	// NewNode is used to broadcast a request to join the graph
	NewNode dispatch.MessageType = "new_node"
)
