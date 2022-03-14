package agent

import (
	"encoding/json"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func NatsConnect(url string) *nats.Conn {
	log.Printf("INFO connecting to NATS at %s", url)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal().
			Str("url", url).
			Err(err).
			Msg("Failed to connect to nats")
	}
	return nc
}

func Subscribe(nc *nats.Conn) {
	log.Info().
		Msg("INFO Subscribing to all subjects")
	subscribeOnline(nc)
	subscribeRebootRequired(nc)
	subscribeUpdatesAvailable(nc)
	subscribeReboot(nc)
	subscribeUpgrade(nc)

	SignalHandler()
}

func subscribeOnline(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.online", func(m *nats.Msg) {
		log.Info().
			Msg("Recieved online query")
		publishOnline(nc)
	})
}

func subscribeUpdatesAvailable(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.updatesavailable", func(m *nats.Msg) {
		log.Info().
			Msg("Recieved updates query")
		publishUpdatesAvailable(nc)
	})
}

func subscribeRebootRequired(nc *nats.Conn) {
	nc.Subscribe("updatemgr.q.rebootrequired", func(m *nats.Msg) {
		log.Info().
			Msg("Recieved reboot query")
		publishRebootRequired(nc)
	})
}

func subscribeUpgrade(nc *nats.Conn) {
	nc.Subscribe("updatemgr.a.upgrade", func(m *nats.Msg) {
		hostname, _ := os.Hostname()
		var res []string
		json.Unmarshal(m.Data, &res)

		for _, i := range res {
			if i == hostname || i == "*" {
				log.Info().
					Str("hostname", i).
					Msg("Recieved upgrade command")
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
				log.Info().
					Str("hostname", i).
					Msg("Recieved reboot command")
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
	log.Info().
		Msg("Publishing online")
	nc.Publish("updatemgr.r.online", []byte(online.Marshall()))
}

func publishUpdatesAvailable(nc *nats.Conn) {
	ua := GetUpdatesAvailable()
	log.Info().
		Msg("Publishing updates")
	nc.Publish("updatemgr.r.updatesavailable", []byte(ua.Marshall()))
}

func publishRebootRequired(nc *nats.Conn) {
	rr := IsRebootRequired()
	log.Info().
		Msg("Publishing reboot required")
	nc.Publish("updatemgr.r.rebootrequired", []byte(rr.Marshall()))
}
