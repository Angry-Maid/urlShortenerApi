package main

import "github.com/go-redis/redis"

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 1,
})

func rSet(key string, val string) bool {
	err := client.Set(key, val, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func rGet(key string) string {
	b, err := client.Exists(key).Result()
	if err != nil {
		return ""
	}
	if b == 1 {
		val, err := client.Get(key).Result()
		if err != nil {
			return ""
		}
		return val
	}
	return ""
}
