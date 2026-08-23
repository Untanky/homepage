package blog

import (
	"context"
)

type Repository interface {
	GetBlog(ctx context.Context, blogID BlogID) (Blog, error)
	GetBlogByHost(ctx context.Context, host string) (Blog, error)
}

type CacheRepository struct {
	hostLookup map[string]BlogID
	cache      map[BlogID]Blog
	origin     Repository
}

func NewBlogRepositoryCache(origin Repository) *CacheRepository {
	return new(CacheRepository{
		hostLookup: make(map[string]BlogID, 16),
		cache:      make(map[BlogID]Blog, 16),
		origin:     origin,
	})
}

func (cache *CacheRepository) GetBlog(ctx context.Context, blogID BlogID) (Blog, error) {
	blg, ok := cache.cache[blogID]
	if ok {
		return blg, nil
	}

	blg, err := cache.origin.GetBlog(ctx, blogID)
	if err != nil {
		return blg, err
	}

	cache.addToCache(blg)

	return blg, nil
}

func (cache *CacheRepository) GetBlogByHost(ctx context.Context, host string) (Blog, error) {
	blogID, ok := cache.hostLookup[host]
	if ok {
		return cache.GetBlog(ctx, blogID)
	}

	blg, err := cache.origin.GetBlogByHost(ctx, host)
	if err != nil {
		return blg, err
	}

	cache.addToCache(blg)

	return blg, nil
}

func (cache *CacheRepository) addToCache(blg Blog) {
	cache.hostLookup[blg.Authority] = blg.ID
	cache.cache[blg.ID] = blg
}
