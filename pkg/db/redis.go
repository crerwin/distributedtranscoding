package db

import (
	"fmt"
	"log"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := new(RedisClient)
	client.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func (c *RedisClient) Ping() {
	pong, err := c.client.Ping().Result()
	fmt.Println(pong, err)
}

func (c *RedisClient) Initialize() {
	log.Print("Deleting the key work_queue")
	err := c.client.Del("work_queue").Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (c *RedisClient) AddToWorkQueue(item string) {
	i := dtc.NewItem("title00.mkv", 1920, 1080)
	err := c.client.LPush("work_queue", i).Err()
	if err != nil {
		fmt.Println(err)
	}
}
