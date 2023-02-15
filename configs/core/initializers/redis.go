package initializers

import (
	"github.com/go-redis/redis"
)

type Cache interface {
	Client() *CacheClient
	Connect() error
	Disconnect() error
}

type CacheClient struct {
	*redis.Client
}

func NewCacheClient(redis *redis.Client) *CacheClient {
	return &CacheClient{redis}
}

type Redis struct {
	// options *redis.Options
	client *CacheClient
}

var _ Cache = (*Redis)(nil)

func NewRedis(client *CacheClient) *Redis {
	return &Redis{client}
}

func (r *Redis) Client() *CacheClient {
	return r.client
}

func (r *Redis) Connect() error {
	return r.client.Client.Ping().Err()
}

func (r *Redis) Disconnect() error {
	return r.client.Close()
}
