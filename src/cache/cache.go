package cache

type Cache interface {
	Save(key, value string) error
	Retrieve(key string) (string, error)
}
