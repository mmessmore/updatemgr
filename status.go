package main

import (
	"encoding/json"
	"fmt"
)

type Online struct {
	HostName  string `json:"hostname"`
	TimeStamp int64  `json:"timestamp"`
}

func UnmarshallOnline(in []byte) *Online {
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
