package dummy

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func NatsConnect(url string) *nats.Conn {
	log.Printf("INFO connecting to NATS at %s", url)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Panicf("FATAL Failed to connect to nats: %v", err)
	}

	log.Print("INFO Subscribing")
	sub, err := nc.SubscribeSync("updatemgr.q.online")
	if err != nil {
		log.Panicf("Error subscribing: %v", err)
	}
	d, _ := time.ParseDuration("1h")
	m, err := sub.NextMsg(d)
	if err != nil {
		log.Panicf("Error getting message: %v", err)
	}
	log.Printf("Got message %b", m.Data)
	// nc.Subscribe("updatemgr.q.online", func(m *nats.Msg) {
	// 	log.Print("INFO Recieved online query")
	// })
	nc.Drain()
	return nc
}
