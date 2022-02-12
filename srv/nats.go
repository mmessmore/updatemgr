package srv

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NatsConnect(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	log.Printf("INFO connecting to NATS at %s", url)
	if err != nil {
		log.Panicf("FATAL Failed to connect to nats: %v", err)
	}
	return nc
}

func PublishQueries(nc *nats.Conn) {
	publishOnline(nc)
	publishUpdatesAvailable(nc)
	publishRebootRequired(nc)
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

func Subscribe(nc *nats.Conn) {
	subscribeOnline(nc)
	subscribeUpdatesAvailable(nc)
	subscribeRebootRequired(nc)
}

func subscribeOnline(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.online", func(m *nats.Msg) {
		online := UnmarshallOnline(m.Data)
		hosts.addOnline(online)
	})
}

func subscribeUpdatesAvailable(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.updates_avilable", func(m *nats.Msg) {
		updatesAvailable := UnmarshallUpdatesAvailable(m.Data)
		hosts.addUpdatesAvailable(updatesAvailable)
	})
}

func subscribeRebootRequired(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.reboot_required", func(m *nats.Msg) {
		rebootRequired := UnmarshallRebootRequired(m.Data)
		hosts.addRebootRequired(rebootRequired)
	})
}
