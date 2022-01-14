package srv

import (
	"encoding/json"
	"fmt"
	"time"
)

type Host struct {
	Name             string
	Online           Online             `json:"online"`
	UpdatesAvailable []UpdatesAvailable `json:"updates_available"`
	RebootRequired   bool
}

func PurgeHosts(hosts []Host, ttl int) {
	oldest := time.Now().Unix() - int64(ttl)

	n := 0
	for _, h := range hosts {
		if h.Online.TimeStamp < oldest {
			hosts[n] = h
			n++
		}
	}
}

/*
	Track which hosts have phoned home recently
*/
type Online struct {
	TimeStamp int64 `json:"timestamp"`
}

func (o Online) Expired(oldest int64) bool {
	return o.TimeStamp < oldest
}

func (o Online) Marshall() string {
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding Online Object")
	}
	return string(out)
}

func Unmarshall(in []byte) *Online {
	o := Online{}
	json.Unmarshal(in, &o)
	return &o
}

/*
	Track what packages are available for update
*/
type UpdatesAvailable struct {
	Packages []string `json:"packages"`
}

func (u UpdatesAvailable) Marshall() string {
	out, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error encoding UpdatesAvailable Object")
	}
	return string(out)
}

func UnmarshallUpdatesAvailable(in []byte) *UpdatesAvailable {
	o := UpdatesAvailable{}
	json.Unmarshal(in, &o)
	return &o
}
