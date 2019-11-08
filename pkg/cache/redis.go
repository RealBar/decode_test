package cache

import (
	"decode_test/pkg/config"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var client *redis.Client

func Setup() {
	client = redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Cfg.RedisConfig.Host + ":" + strconv.Itoa(config.Cfg.RedisConfig.Port),
		Password:     config.Cfg.RedisConfig.Password,
		DialTimeout:  3000,
		PoolSize:     config.Cfg.RedisConfig.Pool.MaxSize,
		MinIdleConns: config.Cfg.RedisConfig.Pool.InitSize,
		IdleTimeout:  config.Cfg.RedisConfig.Pool.IdleTimeOut * time.Millisecond,
	})
}

func Set(key string, value interface{}, expire time.Duration) error {
	return client.Set(key, value, expire).Err()
}

func SetNx(key string, value interface{}, expire time.Duration) error {
	return client.SetNX(key, value, expire).Err()
}

func Delete(key string) error {
	return client.Del(key).Err()
}

func IncrBy(key string, value int) error {
	return client.IncrBy(key, int64(value)).Err()
}

func DecrBy(key string, value int) error {
	return client.DecrBy(key, int64(value)).Err()
}
