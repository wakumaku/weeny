package config

import (
	"os"
	"weeny/cache"
	"weeny/hasher"
)

func CreateCacheFromConfig() cache.Cache {

	// selects cache engine
	if cacheEngine := os.Getenv("CACHE_ENGINE"); cacheEngine == "redis" {
		return cache.NewRedis("redis", 6739, "", 0)
	}

	return cache.NewInMemory()
}

func CreateHasherFromConfig() hasher.Hash {
	return &hasher.Md5{}
}
