package srv

import "github.com/rs/zerolog"

var hosts Store

func RunServer(
	port int,
	ttl int,
	purge int,
	refresh int,
	natsUrl string,
	debug bool,
) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	hosts = InitBoltStore("./updatemgr.db")

	nc := NatsConnect(natsUrl)
	// TODO: this is broken
	// Loops over and over after first run
	go ScheduleAll(purge, ttl, refresh, nc)
	Subscribe(nc)

	RunWebServer(port, nc)
}
