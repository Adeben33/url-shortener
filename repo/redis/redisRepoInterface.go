package redis

import "time"

type redisInterface interface {
	Get(key string) (*string, error)
	Set(key string, value interface{}, expiration time.Duration) (*string, error)
	Delete(key string) error
	KeyRename(key string, newName string) (*string, error)
	PushElementleft(key string, value interface{}) error
	PushElementRight(key string, value interface{}) error
	ListAllElement(key string, start, stop int64) ([]string, error)
	SetAdd(key string, values interface{}) error
	SetHash(key string, values interface{}) error
	GetHash(key string, values []string) ([]interface{}, error)
	GetAllHash(key string) (map[string]string, error)
	DeleteHashfield(key string, fields []string) error
	HashfieldExists(key string, field string) error
}
