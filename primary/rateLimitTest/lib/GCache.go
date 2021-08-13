package lib

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type cacheData struct {
	key      string
	value    interface{}
	expireAt time.Time
}

func newCacheData(key string, value interface{}) *cacheData {
	return &cacheData{key: key, value: value}
}

type GCache struct {
	maxSize int //限制个数 0代表不限制
	elist   *list.List
	edata   map[string]*list.Element
	lock    sync.Mutex
}
type GCacheOpt func(g *GCache)
type GCacheOpts []GCacheOpt

func (this GCacheOpts) apply(g *GCache) {
	for _, fn := range this {
		fn(g)
	}
}

func WithMaxSize(maxSize int) GCacheOpt {
	return func(g *GCache) {
		if maxSize > 0 {
			g.maxSize = maxSize
		}
	}
}

func NewGCache(opts ...GCacheOpt) *GCache {
	cache := &GCache{elist: list.New(), edata: make(map[string]*list.Element)}
	GCacheOpts(opts).apply(cache)
	cache.clear()
	return cache
}

func (this *GCache) Get(key string) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if v, ok := this.edata[key]; ok {
		if !v.Value.(*cacheData).expireAt.After(time.Now()) {
			return nil
		}
		this.elist.MoveToFront(v)
		return v.Value.(*cacheData).value
	}
	return nil
}

const NotExpire = 10 * 24 * 365 * time.Hour

func (this *GCache) Set(key string, newV interface{}, ttl time.Duration) {
	this.lock.Lock()
	defer this.lock.Unlock()
	cache := newCacheData(key, newV)
	if ttl <= 0 {
		cache.expireAt = time.Now().Add(NotExpire)
	} else {
		cache.expireAt = time.Now().Add(ttl)
	}

	if v, ok := this.edata[key]; ok {
		v.Value = cache
		this.elist.MoveToFront(v)
	} else {
		this.edata[key] = this.elist.PushFront(cache)
		if this.maxSize > 0 && len(this.edata) > this.maxSize {
			this.removeOldest()
		}
	}
}

func (this *GCache) Print() {
	ele := this.elist.Front()
	for ele != nil {
		//fmt.Println(this.Get(ele.Value.(*cacheData).key))
		func(key string)  {
			fmt.Println(this.Get(key))
		}(ele.Value.(*cacheData).key)
		ele = ele.Next()
	}
}

func (this *GCache) removeOldest() {
	back := this.elist.Back()
	if back == nil {
		return
	}
	this.remove(back)
}

func (this *GCache) remove(ele *list.Element) {
	key := ele.Value.(*cacheData).key
	delete(this.edata, key)
	this.elist.Remove(ele)
}

func (this *GCache) Len() int {
	return len(this.edata)
}

func (this *GCache) removeExpired() {
	for _, v := range this.edata {
		if !v.Value.(*cacheData).expireAt.After(time.Now()) {
			this.remove(v)
		}
	}
}

func (this *GCache) clear() {
	go func() {
		for {
			this.removeExpired()
			time.Sleep(1 * time.Second)
		}
	}()
}
