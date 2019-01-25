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

// NewRedisClient returns a redis client object
func NewRedisClient() *RedisClient {
	client := new(RedisClient)
	client.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// Ping pings redis
func (c *RedisClient) Ping() {
	pong, err := c.client.Ping().Result()
	fmt.Println(pong, err)
}

// Initialize clears the work queue
func (c *RedisClient) Initialize() {
	log.Print("Flushing database")
	c.client.FlushAll()

}

// AddToWorkQueue adds a new work item to the work queue
func (c *RedisClient) AddToWorkQueue(item *dtc.Item) {
	c.client.Set(item.InputFile+":OutputFile", item.OutputFile, 0)
	c.client.Set(item.InputFile+":Crop", item.Crop, 0)
	err := c.client.LPush("work_queue", item.InputFile).Err()
	if err != nil {
		fmt.Println(err)
	}
}
