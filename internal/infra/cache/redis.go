package cache

import (
	"errors"
	"fmt"
	"sync"

	"go-starter/internal/config"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	once   sync.Once
)

func Setup(conf *config.RedisConfig) error {
	once.Do(func() {
		c := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Password: conf.Password,
			DB:       0,
		})
		if _, err := c.Ping().Result(); err != nil {
			return
		}
		client = c
	})
	if client == nil {
		return errors.New("init redis err")
	}
	return nil
}

func Get() *redis.Client {
	return client
}
