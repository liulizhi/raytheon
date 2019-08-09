package utils

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	redisMaxIdle        = 3
	redisIdleTimeoutSec = 240
)

// NewRedisPool new redis pool
func NewRedisPool() *redis.Pool {
	redisURL := fmt.Sprintf("redis://%s:%d", APIConfig.RedisConfig.Host, APIConfig.RedisConfig.Port)
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connect error: %s", err)
			}
			if APIConfig.RedisConfig.Password != "" {
				if _, authErr := c.Do("AUTH", APIConfig.RedisConfig.Password); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", err)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
