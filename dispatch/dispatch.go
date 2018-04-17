package dispatch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SCKelemen/polaris/graph"
)

func Dispatch(vertex graph.Vertex, message Message) error {
	endpoint := fmt.Sprintf("http://%s/notify", vertex.Address)
	json, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, _ := http.Post(endpoint, "application/json; charset=utf-8", bytes.NewBuffer(json))
	if response.Status != string(http.StatusOK) {
		return Non200Response
	}
	return nil
}

type Exception string

func (e Exception) Error() string { return string(e) }

const (
	Non200Response Exception = "Non200Response"
)
