package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type db struct {
	client *redis.Client
	ctx    context.Context
}

func (d *db) Get(key string) (*string, error) {
	result, err := d.client.Get(d.ctx, key).Result()
	if err != nil {
		log.Println("err getting the key value, %v", err.Error())
		return nil, err
	}
	return &result, nil
}

func (d *db) Set(key string, value interface{}, expiration time.Duration) (*string, error) {
	result, err := d.client.Set(d.ctx, key, value, expiration).Result()
	if err != nil {
		log.Println("err setting the key value, %v", err.Error())
		return nil, err
	}
	return &result, nil
}

func (d *db) Delete(key string) error {
	_, err := d.client.Del(d.ctx, key).Result()
	if err != nil {
		log.Println("err deleting the key value, %v", err.Error())
		return err
	}
	return nil
}

//rename a key name
func (d *db) KeyRename(key string, newName string) (*string, error) {
	result, err := d.client.Rename(d.ctx, key, newName).Result()
	if err != nil {
		log.Println("err getting the key value, %v", err.Error())
		return nil, err
	}
	return &result, nil
}

//PushElementleft this push element to the left of a list
func (d *db) PushElementleft(key string, value interface{}) error {
	_, err := d.client.LPush(d.ctx, key, value).Result()
	if err != nil {
		log.Println("err getting the key value, %v", err.Error())
		return err
	}
	return nil
}

//pushelementright this push element to the left of a list
func (d *db) PushElementRight(key string, value interface{}) error {
	_, err := d.client.RPush(d.ctx, key, value).Result()
	if err != nil {
		log.Println("err getting the key value, %v", err.Error())
		return err
	}
	return nil
}

// Get  elements of a list
func (d *db) ListAllElement(key string, start, stop int64) ([]string, error) {
	result, err := d.client.LRange(d.ctx, key, start, stop).Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

// Add elements to a set
func (d *db) SetAdd(key string, values interface{}) error {
	err := d.client.SAdd(d.ctx, key, values).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return nil
}

//SetMap this is user to save map in the hash store
func (d *db) SetHash(key string, values interface{}) error {
	return d.client.HMSet(d.ctx, key, values).Err()
}

//GetHash this gets the hash/map of in the hash store
func (d *db) GetHash(key string, values []string) ([]interface{}, error) {
	result, err := d.client.HMGet(d.ctx, key, values...).Result()
	if err != nil {
		log.Println("error getting hass %v", err)
		return nil, err
	}
	return result, nil
}

//GetHash this gets the hash/map of in the hash store
func (d *db) GetAllHash(key string) (map[string]string, error) {
	result, err := d.client.HGetAll(d.ctx, key).Result()
	if err != nil {
		log.Println("error getting hass %v", err)
		return nil, err
	}
	return result, nil
}

//DeleteHashfield this delete fields in a hash
func (d *db) DeleteHashfield(key string, fields []string) error {
	return d.client.HDel(d.ctx, key, fields...).Err()
}
