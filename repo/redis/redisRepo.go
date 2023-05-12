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
