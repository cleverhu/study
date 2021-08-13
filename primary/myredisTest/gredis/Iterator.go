package gredis

import "errors"

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

func (this *Iterator) HasNext() bool {
	if this.data == nil || len(this.data) <= this.index {
		return false
	}
	return true
}

func (this *Iterator) Next() interface{} {
	if this.index < len(this.data) {
		defer func() {
			this.index++
		}()
		return this.data[this.index]
	} else {
		panic(errors.New("index range out"))
	}

}
