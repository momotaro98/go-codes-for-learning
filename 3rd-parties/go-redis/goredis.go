package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type goRedisClient struct {
	cli *redis.Client
}

func newGoRedisCli(addr string) *goRedisClient {
	cli := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &goRedisClient{
		cli: cli,
	}
}

func (c *goRedisClient) Close() error {
	return c.cli.Close()
}

func (c *goRedisClient) HSET(key string, obj interface{}) error {
	// _, err := conn.Do("HSET", redis.Args{}.Add(key).AddFlat(obj)...)
	// エラーになる
	_, err := c.cli.HSet(context.Background(), key, obj).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *goRedisClient) HGETALL(key string, dest interface{}) error {
	res := c.cli.HGetAll(context.Background(), key)

	err := res.Scan(dest)
	if err != nil {
		return err
	}

	return nil
}
