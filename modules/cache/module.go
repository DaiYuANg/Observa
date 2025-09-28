package cache

import (
	"context"

	"github.com/eko/gocache/lib/v4/cache"
	redisstore "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("cache_module",
	fx.Provide(newContext, newCache),
)

func newContext() context.Context {
	return context.Background()
}

func newCache(ctx context.Context) *cache.Cache[string] {
	redisStore := redisstore.NewRedis(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}))

	return cache.New[string](redisStore)
}
