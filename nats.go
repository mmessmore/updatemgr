package main

import (
	"github.com/nats-io/nats.go"
	"github.com/prprprus/scheduler"
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

func ScheduledPublishers() {
	s, _ := scheduler.NewScheduler(1000)

	s.Delay().Minute(3).Do(publishOnline)
	s.Delay().Minute(3).Do(publishUpdatesAvailable)
	s.Delay().Minute(3).Do(publishRebootRequired)
}
