/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package srv

var hosts Store

type ServerConfig struct {
	NatsUrl string
	Port    int
	Purge   int
	Refresh int
	Ttl     int
}

func (c *ServerConfig) RunServer() {
	hosts = InitBoltStore("./updatemgr.db")

	nc := NatsConnect(c.NatsUrl)
	// TODO: this is broken
	// Loops over and over after first run
	go ScheduleAll(c.Purge, c.Port, c.Refresh, nc)
	Subscribe(nc)

	RunWebServer(c.Port, nc)
}
