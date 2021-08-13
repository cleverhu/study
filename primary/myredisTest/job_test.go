package main

import (
	"fmt"
	"jtthinkStudy/myredis/redisLocker"
	"sync"
	"testing"
)

var lock *redisLocker.Locker

func init() {
	lock = redisLocker.NewLocker("job")
}

func TestName(t *testing.T) {
	a := 0
	wg := &sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			testJob(&a)
		}()
	}
	wg.Wait()
	fmt.Println(a)
}

func testJob(i *int) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	lock.Lock()
	defer lock.UnLock()
	for j := 1; j <= 100; j++ {
		*i += j
	}
}
