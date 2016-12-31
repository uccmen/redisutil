package redisutil

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisInstance struct {
	pool     *redis.Pool
	host     string
	port     int
	password string // not implemented
}

func NewRedis() *RedisInstance {
	if os.Getenv("REDIS_HOST") == "" {
		log.Fatalln("REDIS_HOST not set")
	}
	if os.Getenv("REDIS_PORT") == "" {
		log.Fatalln("REDIS_PORT not set")
	}

	r := &RedisInstance{}

	r.host = os.Getenv("REDIS_HOST")
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatalln("REDIS_PORT is not an integer")
	}
	r.port = port

	r.pool = r.newPool()
	return r
}

func (r *RedisInstance) DB() *redis.Pool {
	return r.pool
}

func (r *RedisInstance) newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", r.host, r.port))
			if err != nil {
				return nil, err
			}
			if r.password != "" {
				if _, err := c.Do("AUTH", r.password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
