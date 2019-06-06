package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(host string, port int, password string, database int) (Cache, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password,
			DB:       database,
		})

	return &Redis{client}, nil
}

func (r *Redis) Save(key, value string) error {
	_, err := r.client.HSet("urlmaps", key, value).Result()
	return err
}

func (r *Redis) Retrieve(key string) (string, error) {
	return r.client.HGet("urlmaps", key).Result()
}
