package config

import (
	"os"
	"weeny/cache"
	"weeny/hasher"
)

func CreateCacheFromConfig() cache.Cache {

	// selects cache engine
	if cacheEngine := os.Getenv("CACHE_ENGINE"); cacheEngine == "redis" {
		return cache.NewRedis("redis", 6379, "", 0)
	}

	return cache.NewInMemory()
}

func CreateHasherFromConfig() hasher.Hash {
	// selects cache engine
	if enc := os.Getenv("HASHER_ENGINE"); enc == "hashids" {
		return &hasher.Hashids{}
	}

	return &hasher.Md5{}
}
