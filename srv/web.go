/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package srv

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func RunWebServer(port int, nc *nats.Conn) {
	listen(port, nc)
}

// Do the following voodoo to embed our static content in the executable

//go:embed updatemgr-web/public/*
var staticFiles embed.FS

type embedFilesystem struct {
	http.FileSystem
}

func (e embedFilesystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFilesystem{
		FileSystem: http.FS(fsys),
	}
}

/*
	listen configures Gin and starts listening
*/
func listen(port int, nc *nats.Conn) {

	r := gin.Default()

	// ignore cors for now
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// use zerolog for Gin too
	r.Use(ginzerolog.Logger("gin"))

	// Simple ping/pong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Dump all the host data as array
	r.GET("/hosts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hosts": hosts.getHosts(),
		})
	})

	// Dump all the host data as a map
	r.GET("/dahosts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hosts": hosts.daHosts(),
		})
	})

	// Just get a list of hostnames
	r.GET("/hostnames", func(c *gin.Context) {
		var names []string = make([]string, 0)
		data := hosts.getHosts()
		for _, host := range data {
			names = append(names, host.Name)
		}
		c.JSON(200, gin.H{
			"hostsnames": names,
		})
	})

	// Get all info for a given host
	r.GET("/host/:host", func(c *gin.Context) {
		host := c.Param("host")
		p, err := hosts.getHost(host)
		if err != nil {
			c.JSON(404, gin.H{
				"host": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"host": p,
			})
		}
	})

	// Get a list of updates for a given host
	r.GET("/upgrades/:host", func(c *gin.Context) {
		host := c.Param("host")
		p, err := hosts.getUpdatesAvailable(host)
		if err != nil {
			c.JSON(404, gin.H{
				"upgrades": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"upgrades": p,
			})
		}
	})

	// Get whether a reboot is required for a given host
	r.GET("/reboot_required/:host", func(c *gin.Context) {
		host := c.Param("host")
		r, err := hosts.getRebootRequired(host)
		if err != nil {
			c.JSON(404, gin.H{
				"reboot_required": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"reboot_required": r,
			})
		}
	})

	// Force a refresh of data for the universe
	r.POST("/refresh", func(c *gin.Context) {
		PublishQueries(nc)
		c.JSON(200, gin.H{
			"hosts": hosts.getHosts(),
		})
	})

	// Do an upgrade on a given host or "*" for all hosts
	r.POST("/upgrade/:host", func(c *gin.Context) {
		host := c.Param("host")
		PublishUpgrade(nc, []string{host})
	})

	// Reboot a given host or "*" for all hosts
	r.POST("/reboot/:host", func(c *gin.Context) {
		host := c.Param("host")
		PublishReboot(nc, []string{host})
	})

	// Serve our embeded static front-end
	r.Use(static.Serve("/", EmbedFolder(staticFiles, "updatemgr-web/public")))

	// Listen
	r.Run(fmt.Sprintf(":%d", port))
}
