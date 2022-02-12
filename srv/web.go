package srv

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int, natsUrl string) {
	listen(port, natsUrl)
}

func listen(port int, natsUrl string) {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hosts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hosts": hosts.getHosts(),
		})
	})
	r.GET("/refresh", func(c *gin.Context) {
		nc := NatsConnect(natsUrl)
		PublishQueries(nc)
		c.JSON(200, gin.H{
			"hosts": hosts.getHosts(),
		})
	})
	r.Run(fmt.Sprintf(":%d", port))
}
