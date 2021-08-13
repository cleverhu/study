package lib

import (
	"sync"
	"time"
)

type Bucket struct {
	cap   int64
	token int64
	lock  sync.Mutex
	rate  int64
	last  int64
}

func NewBucket(cap int64, rate int64) *Bucket {
	if cap <= 0 || rate <= 0 {
		panic("error cap")
	}
	bucket := &Bucket{cap: cap, token: cap, rate: rate}
	//bucket.start()

	return bucket
}

func (this *Bucket) start() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			this.addToken()
		}
	}()
}

func (this *Bucket) addToken() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.token += this.rate
	if this.token > this.cap {
		this.token = this.cap
	}

}

func (this *Bucket) IsAccept() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	now := time.Now().Unix()
	this.token += (now - this.last) * this.rate
	if this.token > this.cap {
		this.token = this.cap
	}
	this.last = now

	if this.token > 0 {
		this.token--
		return true
	}
	return false
}
