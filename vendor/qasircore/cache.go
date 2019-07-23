package qasircore

import (
	"encoding/json"
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

type cacheapp struct {
	codec *cache.Codec
}

func (this *cacheapp) connection(config map[string]interface{}) {
	urlAddress := config["host"].(string) + ":" + config["port"].(string)
	ring := redis.NewClient(&redis.Options{
		Addr: urlAddress,
		DB:   config["database"].(int),
	})

	codec := &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}

	this.codec = codec
}

func (this *cacheapp) Put(key string, content map[string]interface{}, expired time.Duration) {
	this.codec.Set(&cache.Item{
		Key:        key,
		Object:     content,
		Expiration: expired,
	})
}

func (this *cacheapp) Get(key string) map[string]interface{} {
	var data map[string]interface{}
	if err := this.codec.Get(key, &data); err == nil {
		return data
	} else {
		return nil
	}
}

func NewCache(redisConfig map[string]interface{}) *cacheapp {
	var cache cacheapp

	cache.connection(redisConfig)

	return &cache
}
