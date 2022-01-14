package srv

import (
	"github.com/nats-io/nats.go"
)

var nc, _ = nats.Connect(nats.DefaultURL)

func publishOnline() {
	nc.Publish("updatemgr.q.online", []byte(""))
}

func publishUpdatesAvailable() {
	nc.Publish("updatemgr.q.updatesavailable", []byte(""))
}

func publishRebootRequired() {
	nc.Publish("updatemgr.q.rebootrequired", []byte(""))
}
