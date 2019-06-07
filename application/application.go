package application

import (
	"weeny/cache"
	"weeny/hasher"
)

type Application struct {
	c cache.Cache
	e hasher.Hash
}

func New(cache cache.Cache, encoder hasher.Hash) *Application {
	return &Application{c: cache, e: encoder}
}

func (a *Application) Get(key string) (string, error) {
	return a.c.Retrieve(key)
}

func (a *Application) Save(url string) (string, error) {

	key, err := a.e.Encode(url)
	if err != nil {
		return "", err
	}

	if err := a.c.Save(key, url); err != nil {
		return "", err
	}

	return key, nil
}
