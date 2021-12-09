package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		nc.Publish("ping", []byte("pong"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
