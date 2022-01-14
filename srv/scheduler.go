package srv

import (
	"github.com/nats-io/nats.go"
	"github.com/prprprus/scheduler"
)

func ScheduleAll(purgeInterval int, ttl int, nc *nats.Conn) {
	s, _ := scheduler.NewScheduler(1000)
	ScheduledPublishers(s, purgeInterval, nc)
	SchedulePurge(s, purgeInterval, ttl)
}

func ScheduledPublishers(s *scheduler.Scheduler, purgeInterval int, nc *nats.Conn) {
	s.Every().Minute(purgeInterval).Do(publishOnline, nc)
	s.Every().Minute(purgeInterval).Do(publishUpdatesAvailable, nc)
	s.Every().Minute(purgeInterval).Do(publishRebootRequired, nc)
}

func Purge(ttl int) {
	PurgeHosts(hosts, ttl)
}

func SchedulePurge(s *scheduler.Scheduler, purgeInterval int, ttl int) {
	s.Every().Minute(purgeInterval).Do(Purge, ttl)
}
