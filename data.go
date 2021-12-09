package main

import (
	"encoding/json"
	"fmt"
)

type Online struct {
	HostName  string `json:"hostname"`
	TimeStamp int64  `json:"timestamp"`
}

func RedisGetOnline(hostName string) *Online {
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

func (o Online) Marshall() string {
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding Online Object")
	}
	return string(out)
}

type UpdatesAvailable struct {
	HostName  string   `json:"hostname"`
	TimeStamp int64    `json:"timestamp"`
	Packages  []string `json:"packages"`
}

func RedisGetUpdatesAvailable(hostName string) *UpdatesAvailable {
	res, _ := RedisGetBytes(hostName)
	updatesAvailable := UnmarshallUpdatesAvailable(res)
	updatesAvailable.HostName = hostName
	return updatesAvailable
}

func UnmarshallUpdatesAvailable(in []byte) *UpdatesAvailable {
	o := UpdatesAvailable{}
	json.Unmarshal(in, &o)
	return &o
}

func (o UpdatesAvailable) Marshall() string {
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding UpdatesAvailable Object")
	}
	return string(out)
}

type RebootRequired struct {
	HostName  string `json:"hostname"`
	TimeStamp int64  `json:"timestamp"`
	Required  bool   `json:"reqired"`
}

func RedisGetRebootRequired(hostName string) *RebootRequired {
	res, _ := RedisGetBytes(hostName)
	rebootRequired := UnmarshallRebootRequired(res)
	rebootRequired.HostName = hostName
	return rebootRequired
}

func UnmarshallRebootRequired(in []byte) *RebootRequired {
	o := RebootRequired{}
	json.Unmarshal(in, &o)
	return &o
}

func (o RebootRequired) Marshall() string {
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding RebootRequired Object")
	}
	return string(out)
}
