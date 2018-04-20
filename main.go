package main

import (
	"encoding/json"
	"fmt"

	msgs "github.com/SCKelemen/messages"
	"github.com/SCKelemen/polaris/dispatch"
	"github.com/gin-gonic/gin"
)

func main() {
	dispatcher := dispatch.New()

	router := gin.Default()
	router.POST("/graph", func(c *gin.Context) {

		var msg msgs.Message
		c.BindJSON(&msg)

		switch msg.Type {
		case msgs.BirthMessageType:
			var data msgs.BirthData
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				c.Status(400)
			}
			fmt.Printf("BIRTH: %s %s", data.Name, data.Address)
			dispatcher.Dispatch(msg)
			dispatcher.Subscribe(data.Name, data.Address)
			break
		case msgs.DeathMessageType:
			//fmt.Printf("DEATH: %s %s", data.Name, data.Address)
			dispatcher.Dispatch(msg)
			break
		case msgs.SuicideMessageType:
			//fmt.Printf("SUICIDE: %s %s", data.Name, data.Address)
			dispatcher.Dispatch(msg)
			break
		default:
			break
		}
		c.Status(200)
	})
	router.Run(":9999")

}

type entry struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}
