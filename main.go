package main

import (
	"encoding/json"

	"github.com/SCKelemen/polaris/dispatch"
	"github.com/SCKelemen/polaris/graph"
	"github.com/gin-gonic/gin"
	"github.com/sckelemen/polaris/messages"
)

func main() {
	dispatcher := dispatch.New()

	router := gin.Default()
	router.POST("/graph", func(c *gin.Context) {
		var ent entry
		c.BindJSON(&ent)
		g := graph.NewVertex(ent.Name, ent.Address)
		json, _ := json.Marshal(g)
		message := dispatch.Message{Type: messages.NewNode, Data: string(json)}
		dispatcher.Dispatch(message)
		dispatcher.Subscribe(g.ID, g.Address)
		c.Status(200)
	})
	router.Run(":9999")

}

type entry struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}
