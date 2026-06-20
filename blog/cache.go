package blog

import (
	"context"
)

type Repository interface {
	GetBlog(ctx context.Context, blogID BlogID) (Blog, error)
}

type CacheRepository struct {
	cache  map[BlogID]Blog
	origin Repository
}

func NewBlogRepositoryCache(origin Repository) *CacheRepository {
	return new(CacheRepository {
		cache: make(map[BlogID]Blog, 16),
		origin: origin,
	})
}

func (cache *CacheRepository) GetBlog(ctx context.Context, blogID BlogID) (Blog, error) {
	blg, ok := cache.cache[blogID]
	if ok {
		return blg, nil
	}

	return cache.origin.GetBlog(ctx, blogID)
}
