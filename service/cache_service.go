package service

import (
	"github.com/GoodHot/TinyCMS/cache"
	"github.com/GoodHot/TinyCMS/cache/freecache"
	"github.com/GoodHot/TinyCMS/cache/redis"
	"time"
)

type CacheService struct {
	CacheType string `val:"${cache.type}"`
	Address   string `val:"${cache.url}"`
	Password  string `val:"${cache.password}"`
	DB        int    `val:"${cache.db}"`
	CacheSize int    `val:"${cache.cache_size}"` // cache memory size ：MB
	cacher    cache.Cacher
	isInited  bool
}

func (s *CacheService) initCacher() {
	if s.isInited {
		return
	}
	s.isInited = true
	if s.CacheType == "redis" {
		s.cacher = &redis.Redis{
			Address:  s.Address,
			Password: s.Password,
			DB:       s.DB,
		}
	} else {
		s.cacher = &freecache.FreeCache{
			CacheSize: s.CacheSize * 1024 * 1024,
		}
	}
	s.cacher.Initial()
}

func (s *CacheService) Inc(key string, val int) error {
	s.initCacher()
	return s.cacher.Inc(key, val)
}
func (s *CacheService) SetExp(key string, exp time.Duration) error {
	s.initCacher()
	return s.cacher.SetExp(key, exp)
}
func (s *CacheService) Set(key string, val interface{}, exp time.Duration) error {
	s.initCacher()
	return s.cacher.Set(key, val, exp)
}
func (s *CacheService) Get(key string, i interface{}) error {
	s.initCacher()
	return s.cacher.Get(key, i)
}
func (s *CacheService) GetInt(key string) (int, error) {
	s.initCacher()
	return s.GetInt(key)
}
func (s *CacheService) GetStr(key string) (string, error) {
	s.initCacher()
	return s.GetStr(key)
}
func (s *CacheService) Delete(key string) error {
	s.initCacher()
	return s.Delete(key)
}
