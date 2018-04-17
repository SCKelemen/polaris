package main

import (
	"github.com/SCKelemen/polaris/dispatch"
	"github.com/SCKelemen/polaris/graph"
	"github.com/gin-gonic/gin"
)

func main() {
	vertices := []graph.Vertex{}
	dispatcher := Dispatcher{listeners: vertices}
	router := gin.Default()
	r.POST("/graph", func(c *gin.Context) {
		var ent entry
		c.BindJSON(&ent)
		g := graph.NewVertex(ent.Name, ent.Address)

	})
	//r.GET("/graph")
}

type entry struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

type Dispatcher struct {
	listeners []graph.Vertex
}

func (d Dispatcher) Dispatch(message dispatch.Message) {
	for _, node := range d.listeners {
		dispatch.Dispatch(node, message)
	}
}
