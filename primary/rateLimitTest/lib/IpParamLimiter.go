package lib

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

var IpBucketCache *IpBucketCacheStruct
var IpBucketCache2 *GCache

type IpBucketCacheStruct struct {
	data sync.Map
	sync.WaitGroup
}

func init() {
	IpBucketCache = &IpBucketCacheStruct{}
	IpBucketCache2 = NewGCache(WithMaxSize(10000))
}

func IpLimiter(cap int64, rate int64) func(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			ip := c.Request.RemoteAddr
			var bucket *Bucket
			//if v, ok := IpBucketCache.data.Load(ip); ok {
			//	bucket = v.(*Bucket)
			//} else {
			//	bucket = NewBucket(cap, rate)
			//	IpBucketCache.data.Store(ip, bucket)
			//}
			if v := IpBucketCache2.Get(ip); v != nil {
				bucket = v.(*Bucket)
			} else {
				bucket = NewBucket(cap, rate)

				log.Println(time.Now(), "get from cache")
				IpBucketCache2.Set(ip, bucket, time.Second*5)
			}

			if !bucket.IsAccept() {
				c.AbortWithStatusJSON(429, gin.H{"message": "too many request"})
				return
			}
			handlerFunc(c)
		}
	}
}
