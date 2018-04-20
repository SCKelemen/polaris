package dispatch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	msgs "github.com/SCKelemen/messages"
)

// Dispatcher implements the notifier/subscriber pattern
type Dispatcher struct {
	subscriptions map[string]string
}

// New creates a new Dispatcher
func New() Dispatcher {
	return Dispatcher{subscriptions: make(map[string]string)}
}

// Dispatch sends the message to all subscribers
func (d Dispatcher) Dispatch(message msgs.Message) error {
	json, err := json.Marshal(message)
	if err == nil {
		for _, subscriber := range d.subscriptions {
			endpoint := fmt.Sprintf("http://%s/notify", subscriber)
			http.Post(endpoint, "application/json; charset=utf-8", bytes.NewBuffer(json))
		}
	}
	return err
}

// Subscribe adds the subscription to the dispatcher, and returns
// true if the name doesnt already exist. Otherwise returns false.
func (d Dispatcher) Subscribe(name string, address string) bool {
	_, ok := d.subscriptions[name]
	if !ok {
		d.subscriptions[name] = address
	}
	return !ok
}

// Unsubscribe removes the subscription from the dispatcher,
// and returns true if it was removed. Otherwise it returns false.
func (d Dispatcher) Unsubscribe(name string) bool {
	_, ok := d.subscriptions[name]
	if ok {
		delete(d.subscriptions, name)
	}
	return ok
}
