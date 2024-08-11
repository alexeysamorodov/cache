package cache

import (
	"crypto/md5"
	"runtime"
	"sync"
)

type ShardedCache struct {
	shards     []map[string]string
	locks      []sync.RWMutex
	shardCount int
}

func NewShardedCache() *ShardedCache {
	shardCount := runtime.NumCPU()

	shards := make([]map[string]string, shardCount)
	locks := make([]sync.RWMutex, shardCount)

	for i := range shards {
		shards[i] = make(map[string]string)
	}

	return &ShardedCache{
		shards:     shards,
		locks:      locks,
		shardCount: shardCount,
	}
}

func (c *ShardedCache) Get(key string) (string, error) {
	shardIndex := c.getShardIndex(key)
	c.locks[shardIndex].RLock()
	defer c.locks[shardIndex].RUnlock()
	v, ok := c.shards[shardIndex][key]

	if !ok {
		return "", ErrorNotFound
	}

	return v, nil
}

func (c *ShardedCache) Set(k string, v string) error {
	shardIndex := c.getShardIndex(k)
	c.locks[shardIndex].Lock()
	defer c.locks[shardIndex].Unlock()
	c.shards[shardIndex][k] = v

	return nil
}

func (c *ShardedCache) Delete(k string) error {
	shardIndex := c.getShardIndex(k)
	c.locks[shardIndex].Lock()
	defer c.locks[shardIndex].Unlock()

	delete(c.shards[shardIndex], k)

	return nil
}

func (c *ShardedCache) getShardIndex(key string) int {
	hash := md5.Sum([]byte(key))
	return int(hash[0]) % c.shardCount
}
