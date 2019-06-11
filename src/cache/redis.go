package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(host string, port int, password string, database int) Cache {
	return &Redis{
		redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", host, port),
				Password: password,
				DB:       database,
			}),
	}
}

func (r *Redis) Save(key, value string) error {
	_, err := r.client.HSet("urlmaps", key, value).Result()
	return errors.Wrap(err, "HSET redis")
}

func (r *Redis) Retrieve(key string) (string, error) {
	res, err := r.client.HGet("urlmaps", key).Result()
	if res == "" {
		err = errors.Wrap(err, "HGET redis")
	}
	return res, err
}
