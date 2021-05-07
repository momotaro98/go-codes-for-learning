package main

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

type redigoClient struct {
	pool *redigo.Pool
}

func newRedigoCli(addr string) *redigoClient {
	pool := &redigo.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redigo.Conn, error) { return redigo.Dial("tcp", addr) },
	}
	return &redigoClient{
		pool: pool,
	}
}

func (c *redigoClient) Close() error {
	return c.pool.Close()
}

func (c *redigoClient) HSET(key string, obj interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", redigo.Args{}.Add(key).AddFlat(obj)...)
	if err != nil {
		return err
	}

	return nil
}

func (c *redigoClient) HGETALL(key string, dest interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()

	val, err := redigo.Values(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	err = redigo.ScanStruct(val, dest)
	if err != nil {
		return err
	}

	return nil
}
