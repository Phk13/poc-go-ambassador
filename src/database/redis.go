package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(5 * time.Second)

			key := <-ch
			Cache.Del(context.Background(), key)

			fmt.Printf("Cache %s cleared\n", key)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
