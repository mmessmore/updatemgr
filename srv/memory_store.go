package srv

import (
	"errors"
	"time"
)

type MemoryStore struct {
	Hosts map[string]Host `json:"hosts"`
}

func InitMemoryStore() *MemoryStore {
	hostMap := make(map[string]Host)
	ms := MemoryStore{hostMap}
	return &ms
}

func (m *MemoryStore) addOnline(o Online) {
	h, contains := m.Hosts[o.Name]
	if contains {
		h.Online = o
		m.Hosts[o.Name] = h
	} else {
		newHost := Host{
			Name:   o.Name,
			Online: o,
		}
		m.Hosts[o.Name] = newHost
	}
}

func (m *MemoryStore) addUpdatesAvailable(u UpdatesAvailable) {
	h, contains := m.Hosts[u.Name]
	if contains {
		h.UpdatesAvailable = u
		m.Hosts[u.Name] = h
	} else {
		newHost := Host{
			Name:             u.Name,
			UpdatesAvailable: u,
		}
		m.Hosts[u.Name] = newHost
	}
}

func (m *MemoryStore) addRebootRequired(r RebootRequired) {
	h, contains := m.Hosts[r.Name]
	if contains {
		h.RebootRequired = r
		m.Hosts[r.Name] = h
	} else {
		newHost := Host{
			Name:           r.Name,
			RebootRequired: r,
		}
		m.Hosts[r.Name] = newHost
	}
}
func (m *MemoryStore) getHosts() []Host {
	var hosts []Host = make([]Host, 0)
	for key := range m.Hosts {
		hosts = append(hosts, m.Hosts[key])
	}
	return hosts
}
func (m *MemoryStore) daHosts() map[string]Host {
	return m.Hosts
}

func (m *MemoryStore) getHost(hostname string) (Host, error) {
	val, contains := m.Hosts[hostname]
	if !contains {
		return Host{}, errors.New("No such hostname")
	}
	return val, nil
}

func (m *MemoryStore) getUpdatesAvailable(hostname string) ([]string, error) {
	val, contains := m.Hosts[hostname]
	if !contains {
		return []string{}, errors.New("No such hostname")
	}
	return val.UpdatesAvailable.Packages, nil
}

func (m *MemoryStore) getRebootRequired(hostname string) (bool, error) {
	val, contains := m.Hosts[hostname]
	if !contains {
		return false, errors.New("No such hostname")
	}
	return val.RebootRequired.RebootRequired, nil
}

func (m *MemoryStore) purge(ttl int) {
	oldest := time.Now().Unix() - int64(ttl)

	for k, h := range m.Hosts {
		if h.Online.TimeStamp >= oldest {
			delete(m.Hosts, k)
		}
	}
}
