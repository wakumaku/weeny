package config

import (
	"os"
	"testing"
	"weeny/cache"
	"weeny/hasher"
)

func TestShouldReturnAnInMemoryCacheInstance(t *testing.T) {

	c := CreateCacheFromConfig()

	switch c.(type) {
	case *cache.InMemory:
		// Nice!
	default:
		t.Errorf("Expected *cache.InMemory, got: %T", c)
	}
}

func TestShouldReturnARedisCacheInstance(t *testing.T) {

	os.Setenv("USE_REDIS", "true")
	c := CreateCacheFromConfig()

	switch c.(type) {
	case *cache.Redis:
		// Nice!
	default:
		t.Errorf("Expected *cache.Redis, got: %T", c)
	}
}

func TestShouldReturnAnMD5HasherInstance(t *testing.T) {

	h := CreateHasherFromConfig()

	switch h.(type) {
	case *hasher.Md5:
		// Nice!
	default:
		t.Errorf("Expected *hasher.Md5, got: %T", h)
	}
}
