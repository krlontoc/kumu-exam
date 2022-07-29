package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var cch *cache.Cache

func InitCache() {
	cch = cache.New(2*time.Minute, 1*time.Minute)
}

func GetFromCache(key string) (interface{}, bool) {
	return cch.Get(key)
}

func AddToCache(key string, value interface{}, exp *time.Duration) error {
	itemExp := cache.DefaultExpiration
	if exp != nil {
		itemExp = *exp
	}

	return cch.Add(key, value, itemExp)
}
