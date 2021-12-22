package main

import (
	"encoding/json"
	"fmt"
)


/*
	Track which hosts have phoned home recently
*/
type Online struct {
	HostName  string `json:"hostname"`
	TimeStamp int64  `json:"timestamp"`
}

func (o Online) Marshall() string {
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding Online Object")
	}
	return string(out)
}

func (o Online) SetOnline() error {
	key := fmt.Sprintf("%s:online:%s", RedisKeyPrefix, o.HostName)
	err := RedisSetExInt64(key, RedisDefaultTTl, o.TimeStamp)
	return err
}

func GetOnline(hostName string) *Online {
	o := Online{HostName: hostName}
	ts, _ := RedisGetInt64(hostName)
	o.TimeStamp = ts

	return &o
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
	HostName  string   `json:"hostname"`
	TimeStamp int64    `json:"timestamp"`
	Packages  []string `json:"packages"`
}

func (u UpdatesAvailable) SetUpdatesAvailable() error {
	key := fmt.Sprintf("%s:updates:%s", RedisKeyPrefix, u.HostName)
	err := RedisSetExString(key, RedisDefaultTTl, u.Marshall())
	return err
}

func (u UpdatesAvailable) Marshall() string {
	out, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error encoding UpdatesAvailable Object")
	}
	return string(out)
}
func GetUpdatesAvailable(hostName string) *UpdatesAvailable {
	key := fmt.Sprintf("%s:updates:%s", RedisKeyPrefix, hostName)
	res, _ := RedisGetBytes(key)
	updatesAvailable := UnmarshallUpdatesAvailable(res)
	updatesAvailable.HostName = hostName
	return updatesAvailable
}

func UnmarshallUpdatesAvailable(in []byte) *UpdatesAvailable {
	o := UpdatesAvailable{}
	json.Unmarshal(in, &o)
	return &o
}

/*
	Is a reboot required on a server due to eg a kernel upgrade
*/
type RebootRequired struct {
	HostName  string `json:"hostname"`
	TimeStamp int64  `json:"timestamp"`
	Required  bool   `json:"required"`
}

func (r RebootRequired) SetRebootRequired() error {
	key := fmt.Sprintf("%s:reboot:%s", RedisKeyPrefix, r.HostName)
	err := RedisSetExString(key, RedisDefaultTTl, r.Marshall())
	return err
}

func (r RebootRequired) Marshall() string {
	out, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error encoding RebootRequired Object")
	}
	return string(out)
}

func GetRebootRequired(hostName string) *RebootRequired {
	key := fmt.Sprintf("%s:reboot:%s", RedisKeyPrefix, hostName)
	res, _ := RedisGetBytes(key)
	rebootRequired := UnmarshallRebootRequired(res)
	rebootRequired.HostName = hostName
	return rebootRequired
}

func UnmarshallRebootRequired(in []byte) *RebootRequired {
	o := RebootRequired{}
	json.Unmarshal(in, &o)
	return &o
}

