package srv

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NatsConnect(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Panicf("Failed to connect to nats: %v", err)
	}
	return nc
}

func publishOnline(nc *nats.Conn) {
	nc.Publish("updatemgr.q.online", []byte(""))
}

func publishUpdatesAvailable(nc *nats.Conn) {
	nc.Publish("updatemgr.q.updatesavailable", []byte(""))
}

func publishRebootRequired(nc *nats.Conn) {
	nc.Publish("updatemgr.q.rebootrequired", []byte(""))
}
