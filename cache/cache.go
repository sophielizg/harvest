package cache

type Cache interface {
	Exists(key string) (bool, error)
	Get(key string) (any, error)
	Put(key string, item any, ttl int) error
}
