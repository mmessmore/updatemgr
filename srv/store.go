/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package srv

type Store interface {
	addOnline(Online)
	addUpdatesAvailable(UpdatesAvailable)
	getUpdatesAvailable(string) ([]string, error)
	addRebootRequired(RebootRequired)
	getRebootRequired(string) (bool, error)
	getHosts() []Host
	daHosts() map[string]Host
	getHost(string) (Host, error)
	purge(ttl int)
}
