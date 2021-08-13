package eventbus

import (
	"fmt"
	"sync"
)

type EventBus struct {
	subscribes map[string][]EventDataChan
	handlers   map[string]*EventHandler
	lock       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{subscribes: make(map[string][]EventDataChan), handlers: make(map[string]*EventHandler)}
}

func (this *EventBus) Sub(topic string, fn interface{}) EventDataChan {
	this.lock.Lock()
	defer this.lock.Unlock()
	ec := make(EventDataChan)
	this.subscribes[topic] = append(this.subscribes[topic], ec)
	this.handlers[topic] = NewEventHandler(fn)
	return ec

}

func (this *EventBus) UnSub(topic string, ch EventDataChan) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.subscribes[topic]; ok {

		this.removeSubscribe(topic, ch)
	}
}

func (this *EventBus) removeSubscribe(topic string, ch EventDataChan) {
	chs := this.subscribes[topic]

	for i := 0; i < len(chs); i++ {
		if chs[i] == ch {
			this.subscribes[topic] = append(this.subscribes[topic][:i], this.subscribes[topic][i+1:]...)
		}
	}
}

func (this *EventBus) Pub(topic string, ch EventDataChan, params ...interface{}) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if ecs, found := this.subscribes[topic]; found {
		for _, ec := range ecs {
			if ch != nil && ec == ch {
				go func(inCh EventDataChan) {
					inCh <- &EventData{Data: this.handlers[topic].Call(params...)}
				}(ec)
			}

		}

	}
	//else {
	//	go func() {
	//		this.subscribes[topic] = make(EventDataChan)
	//		this.subscribes[topic] <- &EventData{Data: data}
	//	}()
	//
	//}
}

func (this *EventBus) PrintNum(topic string) {
	fmt.Println(len(this.subscribes[topic]))
}
