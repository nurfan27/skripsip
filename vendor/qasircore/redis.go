package qasircore

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (this *Redis) GetClient() *redis.Client {
	return this.client
}

func (this *Redis) Set(key string, value string) {
	this.client.Set(key, value, 0)
}

func (this *Redis) Get(key string) string {
	val, err := this.client.Get(key).Result()

	if err != nil {
		return ""
	} else {
		return val
	}
}

func (this *Redis) SetupConfiguration(config map[string]interface{}) {
	urlAddress := config["host"].(string) + ":" + config["port"].(string)

	this.client = redis.NewClient(&redis.Options{
		Addr:     urlAddress,
		Password: config["password"].(string),
		DB:       config["db"].(int),
	})
}

func RedisClient(config map[string]interface{}) *Redis {
	var redisclient Redis

	redisclient.SetupConfiguration(config)

	return &redisclient
}
