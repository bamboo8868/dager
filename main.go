package main

import (
	"github.com/dager/redis"
	"log"
	"time"
)

func main() {
	config := redis.Config{
		Port: 6379,
		Addr: []byte("127.0.0.1"),
		Pass: []byte(""),
		DB:   0,
	}
	client := redis.New(&config)

	res := client.Get("name")
	if res {
		log.Println("success")
	}else {
		log.Println("error")
	}
	time.Sleep(1*time.Second)
}
