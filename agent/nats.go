package agent

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func NatsConnect(url string) *nats.Conn {
	log.Printf("INFO connecting to NATS at %s", url)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Panicf("FATAL Failed to connect to nats: %v", err)
	}
	return nc
}

func Subscribe(nc *nats.Conn) {
	subscribeOnline(nc)
	subscribeRebootRequired(nc)
	subscribeUpdatesAvailable(nc)
	subscribeReboot(nc)
	subscribeUpgrade(nc)
}

func subscribeOnline(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.online", func(m *nats.Msg) {
		publishOnline(nc)
	})
}

func subscribeUpdatesAvailable(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.updates_available", func(m *nats.Msg) {
		publishUpdatesAvailable(nc)
	})
}

func subscribeRebootRequired(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.updates_available", func(m *nats.Msg) {
		publishRebootRequired(nc)
	})
}

func subscribeUpgrade(nc *nats.Conn) {
	nc.Subscribe("updatemgr.a.update", func(m *nats.Msg) {
		hostname, _ := os.Hostname()
		var res []string
		json.Unmarshal(m.Data, &res)

		for _, i := range res {
			if i == hostname || i == "*" {
				DoUpgrade()
				break
			}
		}
	})
}

func subscribeReboot(nc *nats.Conn) {
	nc.Subscribe("updatemgr.a.reboot", func(m *nats.Msg) {
		hostname, _ := os.Hostname()
		var res []string
		json.Unmarshal(m.Data, &res)

		for _, i := range res {
			if i == hostname || i == "*" {
				DoReboot()
				break
			}
		}
	})
}

func publishOnline(nc *nats.Conn) {
	hostname, _ := os.Hostname()
	online := Online{
		Name:      hostname,
		TimeStamp: time.Now().Unix(),
	}
	nc.Publish("updatemgr.r.online", []byte(online.Marshall()))
}

func publishUpdatesAvailable(nc *nats.Conn) {
	ua := GetUpdatesAvailable()
	nc.Publish("updatemgr.r.updatesavailable", []byte(ua.Marshall()))
}

func publishRebootRequired(nc *nats.Conn) {
	rr := IsRebootRequired()
	nc.Publish("updatemgr.r.rebootrequired", []byte(rr.Marshall()))
}
