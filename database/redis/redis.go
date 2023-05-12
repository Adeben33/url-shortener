package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var db *redis.Client

func GetRedisDb() *redis.Client {
	return db
}

func ConnectToRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println("connecting to redis database")
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		log.Panic(err)
		return nil
	}
	log.Printf("redis database pinged %v", ping)
	db = client
	return client
}
