package cache

import (
	"errors"
	"sync"
)

type InMemory struct {
	storage sync.Map
}

func NewInMemory() Cache {
	return &InMemory{}
}

func (r *InMemory) Save(key, value string) error {
	r.storage.Store(key, value)
	return nil
}

func (r *InMemory) Retrieve(key string) (string, error) {
	v, f := r.storage.Load(key)
	if !f {
		return "", errors.New("key not found")
	}

	return v.(string), nil
}
