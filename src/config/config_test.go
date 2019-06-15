package config

import (
	"os"
	"testing"
	"weeny/cache"
	"weeny/hasher"
)

func TestShouldReturnAnInMemoryCacheInstance(t *testing.T) {

	lastVal := os.Getenv("CACHE_ENGINE")
	os.Setenv("CACHE_ENGINE", "")
	defer os.Setenv("CACHE_ENGINE", lastVal)

	c := CreateCacheFromConfig()

	switch c.(type) {
	case *cache.InMemory:
		// Nice!
	default:
		t.Errorf("Expected *cache.InMemory, got: %T", c)
	}
}

func TestShouldReturnARedisCacheInstance(t *testing.T) {

	lastVal := os.Getenv("CACHE_ENGINE")
	os.Setenv("CACHE_ENGINE", "redis")
	defer os.Setenv("CACHE_ENGINE", lastVal)

	c := CreateCacheFromConfig()

	switch c.(type) {
	case *cache.Redis:
		// Nice!
	default:
		t.Errorf("Expected *cache.Redis, got: %T", c)
	}
}

func TestShouldReturnAnMD5HasherInstanceByDefault(t *testing.T) {

	lastVal := os.Getenv("HASHER_ENGINE")
	os.Setenv("HASHER_ENGINE", "")
	defer os.Setenv("HASHER_ENGINE", lastVal)

	h := CreateHasherFromConfig()

	switch h.(type) {
	case *hasher.Md5:
		// Nice!
	default:
		t.Errorf("Expected *hasher.Md5, got: %T", h)
	}
}

func TestShouldReturnAHashidsHasherInstance(t *testing.T) {

	lastVal := os.Getenv("HASHER_ENGINE")
	os.Setenv("HASHER_ENGINE", "hashids")
	defer os.Setenv("HASHER_ENGINE", lastVal)

	h := CreateHasherFromConfig()

	switch h.(type) {
	case *hasher.Hashids:
		// Nice!
	default:
		t.Errorf("Expected *hasher.Hashids, got: %T", h)
	}
}
