package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"nitra/registry/utils/propertymanager"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitializeRedis()  {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     propertymanager.GetStringProperty("redis.host", "localhost") + ":" + propertymanager.GetStringProperty("redis.port", "6379") ,
		Password: propertymanager.GetStringProperty("redis.password", ""),
		DB:       propertymanager.GetIntProperty("redis.database", 0),
		PoolSize: propertymanager.GetIntProperty("redis.conn.pool-size", 10),
	})
	log.Println("RedisClient initialized successfully...")
}

func HSet(key, field, value string) error  {
	return redisClient.HSet(ctx, key, field, value).Err()
}

func HGet(key, field string) (string, error) {
	return redisClient.HGet(ctx, key, field).Result()
}

func HDel(key, field string) (int64, error) {
	return redisClient.HDel(ctx, key, field).Result()
}

func HGetAll(key string)  (map[string]string, error){
	return redisClient.HGetAll(ctx, key).Result()
}

func IsRedisNil(err error) bool  {
	if err == redis.Nil {
		return true
	}
	return false
}