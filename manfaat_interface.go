package main

import "fmt"

type cacheRedisService struct{}

func (cache *cacheRedisService) Get(key string) string {
	return "<cache from redis>"
}

func NewCacheRedisService() *cacheRedisService {
	return &cacheRedisService{}
}

type inMemoryCacheService struct{}

func (cache *inMemoryCacheService) GetCache(key string) string {
	return "<cache from in memory>"
}

func NewInMemoryCacheService() *inMemoryCacheService {
	return &inMemoryCacheService{}
}

func main() {

	// cacheService := NewCacheRedisService()
	cacheService := NewInMemoryCacheService()

	myCache := cacheService.GetCache("my-key")

	fmt.Println(myCache)

}
