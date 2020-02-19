package freecache

import (
	"github.com/coocood/freecache"
	"github.com/vmihailenco/msgpack"
	"time"
)

type FreeCache struct {
	CacheSize int
	cache     *freecache.Cache
}

func (s *FreeCache) Initial() {
	s.cache = freecache.NewCache(s.CacheSize)
}

func (s *FreeCache) Inc(key string, val int) error {
	return nil
}

func (s *FreeCache) SetExp(key string, exp time.Duration) error {
	return nil
}

func (s *FreeCache) Set(key string, val interface{}, exp time.Duration) error {
	data, err := s.marshal(val)
	if err != nil {
		return err
	}
	return s.cache.Set([]byte(key), data, int(exp.Seconds()))
}

func (s *FreeCache) Get(key string, i interface{}) error {
	data, err := s.cache.Get([]byte(key))
	if err != nil {
		return err
	}
	return s.unmarshal(data, i)
}

func (s *FreeCache) GetInt(key string) (int, error) {
	data, err := s.cache.Get([]byte(key))
	if err != nil {
		return -1, err
	}
	var rst int
	err = s.unmarshal(data, rst)
	return rst, err
}

func (s *FreeCache) GetStr(key string) (string, error) {
	data, err := s.cache.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *FreeCache) Delete(key string) error {
	s.cache.Del([]byte(key))
	return nil
}

func (*FreeCache) marshal(i interface{}) ([]byte, error) {
	return msgpack.Marshal(i)
}

func (*FreeCache) unmarshal(bts []byte, i interface{}) error {
	return msgpack.Unmarshal(bts, i)
}
