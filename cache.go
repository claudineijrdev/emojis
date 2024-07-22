package main

import (
	"context"
	"sync"

	"github.com/ServiceWeaver/weaver"
)


type Cache interface {
	Get(context.Context, string) ([]string, error)
	Put(context.Context, string, []string) error
}

var _ weaver.NotRetriable = Cache.Put

type cache struct {
	weaver.Implements[Cache]
	mu sync.Mutex
	emojis map[string][]string
}

func (c *cache) Init(context.Context) error {
	c.emojis = map[string][]string{}
	return nil
}

func (c *cache) Put(ctx context.Context, key string, value []string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger(ctx).Debug("Cache Put","query",value)
	c.emojis[key] = value
	return nil
}

func (c *cache) Get(ctx context.Context, key string) ([]string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger(ctx).Debug("Cache Get","query",key)
	return c.emojis[key], nil
}