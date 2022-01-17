package srv

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) {
	listen(port)
}

func listen(port int) {

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
	r.Run(fmt.Sprintf(":%d", port))
}
