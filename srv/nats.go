package srv

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func NatsConnect(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	log.Info().
		Str("url", url).
		Msg("Connecting to NATS")
	if err != nil {
		log.Fatal().
			Str("url", url).
			Err(err).
			Msg("Failed to connect to nats")
	}
	return nc
}

func PublishQueries(nc *nats.Conn) {
	PublishOnline(nc)
	PublishUpdatesAvailable(nc)
	PublishRebootRequired(nc)
}

func PublishOnline(nc *nats.Conn) {
	log.Debug().
		Msg("Publishing online query")
	nc.Publish("updatemgr.q.online", []byte(""))
}

func PublishUpdatesAvailable(nc *nats.Conn) {
	log.Debug().
		Msg("Publishing update query")
	nc.Publish("updatemgr.q.updatesavailable", []byte(""))
}

func PublishRebootRequired(nc *nats.Conn) {
	log.Debug().
		Msg("Publishing reboot required query")
	nc.Publish("updatemgr.q.rebootrequired", []byte(""))
}

func PublishUpgrade(nc *nats.Conn, hosts []string) {
	log.Info().
		Strs("hosts", hosts).
		Msg("Requesting upgrade for hosts")
	hostsJson, _ := json.Marshal(hosts)
	nc.Publish("updatemgr.a.upgrade", hostsJson)
}

func PublishReboot(nc *nats.Conn, hosts []string) {
	log.Info().
		Strs("hosts", hosts).
		Msg("Requesting reboot for hosts")
	hostsJson, _ := json.Marshal(hosts)
	nc.Publish("updatemgr.a.reboot", hostsJson)
}

func Subscribe(nc *nats.Conn) {
	log.Debug().
		Msg("Subscribing to all")
	subscribeOnline(nc)
	subscribeUpdatesAvailable(nc)
	subscribeRebootRequired(nc)
}

func subscribeOnline(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.online", func(m *nats.Msg) {
		online := UnmarshallOnline(m.Data)
		log.Debug().
			Str("host", online.Name).
			Msg("Got online response")
		hosts.addOnline(*online)
	})
}

func subscribeUpdatesAvailable(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.updatesavailable", func(m *nats.Msg) {
		log.Printf("INFO: got updates response:\n%s\n", string(m.Data))
		updatesAvailable := UnmarshallUpdatesAvailable(m.Data)
		log.Debug().
			Str("host", updatesAvailable.Name).
			Msg("Got updates response")
		hosts.addUpdatesAvailable(*updatesAvailable)
	})
}

func subscribeRebootRequired(nc *nats.Conn) {
	nc.Subscribe("updatemgr.r.rebootrequired", func(m *nats.Msg) {
		log.Printf("INFO: got reboot response:\n%s\n", string(m.Data))
		rebootRequired := UnmarshallRebootRequired(m.Data)
		log.Debug().
			Str("host", rebootRequired.Name).
			Msg("Got reboot response")
		hosts.addRebootRequired(*rebootRequired)
	})
}
