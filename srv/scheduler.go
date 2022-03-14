package srv

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// TODO: this is broken
// Fast loop after initial wait
func ScheduleAll(purgeInterval int, refresh int, ttl int, nc *nats.Conn) {
	go func() {
		for {
			publishAllScheduled(nc)
			time.Sleep(time.Duration(ttl) * time.Second)
		}
	}()
}

func publishAllScheduled(nc *nats.Conn) {
	log.Info().
		Msg("Publishing scheduled queries")
	PublishOnline(nc)
	PublishUpdatesAvailable(nc)
	PublishRebootRequired(nc)
}
