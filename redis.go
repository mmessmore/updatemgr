package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gomodule/redigo/redis"
)

var Pool = newPool()

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func RedisGetBytes(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func RedisGetInt64(key string) (int64, error) {
	conn := Pool.Get()
	defer conn.Close()
	var data int64
	data, err := redis.Int64(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}
