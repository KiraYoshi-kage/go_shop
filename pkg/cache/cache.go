package cache

import "time"

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
}
