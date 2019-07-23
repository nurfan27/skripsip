package interfaces

import (
	"time"
)

type CacheInterface interface {
	Put(key string, content map[string]interface{}, expired time.Duration)
	Get(key string) map[string]interface{}
}
