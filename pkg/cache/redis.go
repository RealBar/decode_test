package cache

import (
	"decode_test/pkg/config"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Cache struct {
	cli *redis.Client
}

func Setup(cfg *config.Config, logger *logrus.Logger) *Cache {
	client2 := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         cfg.RedisConfig.Host + ":" + strconv.Itoa(cfg.RedisConfig.Port),
		Password:     cfg.RedisConfig.Password,
		DialTimeout:  3000,
		PoolSize:     cfg.RedisConfig.Pool.MaxSize,
		MinIdleConns: cfg.RedisConfig.Pool.InitSize,
		IdleTimeout:  cfg.RedisConfig.Pool.IdleTimeOut * time.Millisecond,
	})
	pong, err := client2.Ping().Result()
	if err != nil {
		logger.WithError(err).Error("setup redis client2 error")
		panic(err)
	}
	logger.Info("redis ping result:", pong)
	return &Cache{cli: client2}
}

func (c *Cache) Set(key string, value interface{}, expire time.Duration) error {
	return c.cli.Set(key, value, expire).Err()
}

func (c *Cache) SetNx(key string, value interface{}, expire time.Duration) error {
	return c.cli.SetNX(key, value, expire).Err()
}

func (c *Cache) Delete(key string) error {
	return c.cli.Del(key).Err()
}

func (c *Cache) IncrBy(key string, value int) (int64, error) {
	result := c.cli.IncrBy(key, int64(value))
	return result.Val(), result.Err()
}

func (c *Cache) DecrBy(key string, value int) (int64, error) {
	result := c.cli.DecrBy(key, int64(value))
	return result.Val(), result.Err()
}
