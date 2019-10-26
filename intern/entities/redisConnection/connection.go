// connection
package redisConnection

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func InitPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
			}
			return conn, err
		},
	}
}

func Get() redis.Conn {
	return pool.Get()
}
