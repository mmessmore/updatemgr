/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package srv

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Host struct {
	Name             string           `json:"name"`
	Online           Online           `json:"online"`
	UpdatesAvailable UpdatesAvailable `json:"updates_available"`
	RebootRequired   RebootRequired   `json:"reboot_required"`
}

func (h Host) Marshall() []byte {
	var out = []byte("")
	out, err := json.Marshal(h)
	if err != nil {
		log.Error().
			Err(err).
			Str("host", h.Name).
			Msg("Error marshalling to JSON")
	}

	return out
}

func UnmarshallHost(in []byte) *Host {
	h := Host{}
	json.Unmarshal(in, &h)
	return &h
}

/*
	Track which hosts have phoned home recently
*/
type Online struct {
	Name      string `json:"name"`
	TimeStamp int64  `json:"timestamp"`
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

func UnmarshallOnline(in []byte) *Online {
	o := Online{}
	json.Unmarshal(in, &o)
	return &o
}

/*
	Track what packages are available for update
*/
type UpdatesAvailable struct {
	Name     string   `json:"name"`
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

/*
	Track what hosts require reboot
*/
type RebootRequired struct {
	Name           string `json:"name"`
	RebootRequired bool   `json:"reboot_required"`
}

func (r RebootRequired) Marshall() string {
	out, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error encoding RebootRequired Object")
	}
	return string(out)
}

func UnmarshallRebootRequired(in []byte) *RebootRequired {
	r := RebootRequired{}
	json.Unmarshal(in, &r)
	return &r
}
