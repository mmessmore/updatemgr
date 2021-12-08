package main

import "github.com/gomodule/redigo/redis"

var pool = newPool()

func newPool() *redis.Pool {
     return &redis.Pool{
         MaxIdle: 10,
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
