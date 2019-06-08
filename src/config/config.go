package config

import (
	"os"
	"weeny/cache"
	"weeny/hasher"
)

func CreateCacheFromConfig() cache.Cache {

	// if setted, just use the default paramsfor redis
	if useRedis := os.Getenv("USE_REDIS"); useRedis != "" {
		return cache.NewRedis("redis", 6739, "", 0)
	}

	return cache.NewInMemory()
}

func CreateHasherFromConfig() hasher.Hash {
	return &hasher.Md5{}
}
