package lib

import (
	"jtthinkStudy/myredis/gredis"
	"sync"
	"time"
)

var CachePool *sync.Pool

func init() {
	CachePool = &sync.Pool{New: func() interface{} {
		return gredis.NewSimpleCache(gredis.NewStringOperator(), time.Second*15, gredis.Serialize_JSON)
	}}
}

func NewSimpleCache() *gredis.SimpleCache {
	return CachePool.Get().(*gredis.SimpleCache)
}

func ReleaseSimpleCache(cache *gredis.SimpleCache) {
	CachePool.Put(cache)
}
