package main

import (
	"fmt"
)

type Post struct {
	Title  string `redis:"title"`
	UserID int64  `redis:"user_id"`
}

type RedisCli interface {
	Close() error
	HSET(key string, obj interface{}) error
	HGETALL(key string, dest interface{}) error
}

func run(redisCli RedisCli) error {
	defer redisCli.Close()

	post := &Post{
		Title:  "First Game",
		UserID: 123,
	}

	key := "key1"
	if err := redisCli.HSET(key, post); err != nil {
		return err
	}
	var outPost Post
	if err := redisCli.HGETALL(key, &outPost); err != nil {
		return err
	}

	fmt.Println(outPost)

	return nil
}

func main() {
	addr := "localhost:6379"
	clis := []RedisCli{
		// newRedigoCli(addr),
		newGoRedisCli(addr),
	}
	for _, cli := range clis {
		if err := run(cli); err != nil {
			panic(err)
		}
	}
}
