package srv

var hosts = make([]Host, 5)

func RunServer(port int, ttl int, purge int, natsUrl string) {

	nc := NatsConnect(natsUrl)
	go ScheduleAll(purge, ttl, nc)

	RunWebServer(port)
}
