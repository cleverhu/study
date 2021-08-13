package eventbus

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 事件数据
type EventData struct {
	Data interface{}
}

// 传递事件数据的通道
type EventDataChan chan *EventData

func (this EventDataChan) Data(timeout time.Duration) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return gin.H{"message": "timeout"}
		case data := <-this:
			return data.Data
		}
	}
}
