package redis

import (
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"time"
)

type Redis struct {
	Address  string
	Password string
	DB       int
	client   *redis.Client
}

func (s *Redis) Initial() {
	s.client = redis.NewClient(&redis.Options{
		Addr:     s.Address,
		Password: s.Password,
		DB:       s.DB,
	})
	_, err := s.client.Ping().Result()
	if err != nil {
		// TODO redis log and change freecache
	}
}

func (s *Redis) Inc(key string, val int) error {
	return s.client.IncrBy(key, int64(val)).Err()
}

func (s *Redis) SetExp(key string, exp time.Duration) error {
	return s.client.Expire(key, exp).Err()
}

func (s *Redis) Set(key string, val interface{}, exp time.Duration) error {
	bt, err := s.marshal(val)
	if err != nil {
		return err
	}
	return s.client.Set(key, bt, exp).Err()
}

func (s *Redis) Get(key string, i interface{}) error {
	result := s.client.Get(key)
	if result.Err() != nil {
		return result.Err()
	}
	data, err := result.Bytes()
	if err != nil {
		return err
	}
	return s.unmarshal(data, i)
}

func (s *Redis) GetInt(key string) (int, error) {
	result := s.client.Get(key)
	if result.Err() != nil {
		return -1, result.Err()
	}
	return result.Int()
}

func (s *Redis) GetStr(key string) (string, error) {
	result := s.client.Get(key)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.String(), result.Err()
}

func (s *Redis) Delete(key string) error {
	return s.client.Del(key).Err()
}

func (*Redis) marshal(i interface{}) ([]byte, error) {
	return msgpack.Marshal(i)
}

func (*Redis) unmarshal(bts []byte, i interface{}) error {
	return msgpack.Unmarshal(bts, i)
}
