package srv

type MemoryStore struct {
	Hosts map[string]Host `json:"hosts"`
}

func InitMemoryStore() MemoryStore {
	var hostMap map[string]Host
	ms := MemoryStore{hostMap}
	return ms
}

func (m MemoryStore) addOnline(o *Online) {
	host := m.Hosts[o.Name]
	host.Online = *o
}

func (m MemoryStore) addUpdatesAvailable(u *UpdatesAvailable) {
	host := m.Hosts[u.Name]
	host.UpdatesAvailable = *u
}

func (m MemoryStore) addRebootRequired(r *RebootRequired) {
	host := m.Hosts[r.Name]
	host.RebootRequired = *r
}

func (m MemoryStore) getHosts() []Host {
	var hosts []Host = make([]Host, 0)
	for key := range m.Hosts {
		hosts = append(hosts, m.Hosts[key])
	}
	return hosts
}
