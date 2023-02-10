package cache

import (
	"github.com/patrickmn/go-cache"
	"log"

	"time"
)

var OpenAiClientCache *cache.Cache
var AskRequestLockCache *cache.Cache

func Init() {
	// 默认过期时间10分钟；清理间隔60s，即每60s钟会自动清理过期的键值对
	OpenAiClientCache = cache.New(10*60*time.Second, 60*time.Second)
	// 五分钟超时
	AskRequestLockCache = cache.New(5*60*time.Second, 60*time.Second)
	log.Printf("cache init success.")
}
