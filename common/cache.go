package common

import (
	"fmt"
	"time"
)

type ICache interface {
	Get(key string, i interface{}) error
	Set(key string, i interface{}, ttl ...time.Duration) error
}

type Cache struct {
	Type      string `val:"${props.cache.type}"`
	Freecache *struct {
		CacheSize int `json:"cacheSize"`
	} `val:"${props.cache.freecache}"`
	Redis *struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"datasource"`
	} `val:"${props.cache.redis}"`
	instance ICache
}

func (cache *Cache) Ins() ICache {
	if cache.instance != nil {
		return cache.instance
	}
	if cache.Type == "redis" {
		cache.instance = &RedisCache{}
	} else {
		cache.instance = &FreeCache{}
	}
	return cache.instance
}

/**
Member cache
*/
type FreeCache struct {
}

func (*FreeCache) Get(key string, i interface{}) error {
	fmt.Println(key, i)
	return nil
}

func (*FreeCache) Set(key string, i interface{}, ttl ...time.Duration) error { return nil }

/**
Redis cache
*/
type RedisCache struct {
}

func (*RedisCache) Get(key string, i interface{}) error {
	return nil
}

func (*RedisCache) Set(key string, i interface{}, ttl ...time.Duration) error { return nil }
