package cache

import "sync"

type ConcurrentCache struct {
	items map[string]string
	mutex sync.RWMutex
}

func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		items: make(map[string]string),
	}
}

func (c *ConcurrentCache) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, ok := c.items[key]

	if !ok {
		return "", ErrorNotFound
	}

	return item, nil
}

func (c *ConcurrentCache) Set(key string, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = value

	return nil
}

func (c *ConcurrentCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)

	return nil
}
