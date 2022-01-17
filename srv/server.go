package srv

var hosts Store

func RunServer(port int, ttl int, purge int, natsUrl string) {
	store := InitMemoryStore()
	hosts = store

	nc := NatsConnect(natsUrl)
	go ScheduleAll(purge, ttl, nc)
	Subscribe(nc)

	RunWebServer(port)

	defer nc.Drain()
}
