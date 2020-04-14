package cache

import "time"

type Cacher interface {
	Initial()
	Inc(key string, val int) error
	SetExp(key string, exp time.Duration) error
	Set(key string, val interface{}, exp time.Duration) error
	Get(key string, i interface{}) error
	GetInt(key string) (int, error)
	GetStr(key string) (string, error)
	Delete(key string) error
}