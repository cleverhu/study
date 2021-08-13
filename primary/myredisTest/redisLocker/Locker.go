package redisLocker

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"jtthinkStudy/myredis/gredis"
	"time"
)

type Locker struct {
	redis      *redis.Client
	key        string
	expire     time.Duration
	unLock     bool
	incrScript *redis.Script
}

const incrLua = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

func NewLocker(key string) *Locker {
	return &Locker{key: key, expire: time.Second * 5, redis: gredis.Redis(), incrScript: redis.NewScript(incrLua)}
}

func NewLockerWithTTL(key string, expire time.Duration) *Locker {
	if expire.Seconds() <= 0 {
		panic("error expire")
	}
	return &Locker{key: key, expire: expire, redis: gredis.Redis(), incrScript: redis.NewScript(incrLua)}
}

func (this *Locker) Lock() *Locker {
	ok, err := this.redis.SetNX(context.Background(), this.key, "1", 0).Result()
	if !ok || err != nil {
		panic("lock error")
	}
	this.expireFunc()
	return this
}

func (this *Locker) UnLock() {
	this.unLock = true
	this.redis.Del(context.Background(), this.key)
}

func (this *Locker) expireFunc() {
	sleepTime := 2.0 / 3 * this.expire.Seconds()
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleepTime))
			if this.unLock {
				break
			}
			this.resetExpireTime()
		}
	}()
}

func (this *Locker) resetExpireTime() {
	result, err := this.incrScript.Run(context.Background(), this.redis, []string{this.key}, 1, this.expire.Seconds()).Result()
	fmt.Printf("%s续期结果:%v:%v\n", this.key, result, err)
}
