package srv

type Store interface {
	addOnline(*Online)
	addUpdatesAvailable(*UpdatesAvailable)
	addRebootRequired(*RebootRequired)
	getHosts() []Host
}
