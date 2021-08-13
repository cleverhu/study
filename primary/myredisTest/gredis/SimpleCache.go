package gredis

import (
	"encoding/json"
	"time"
)

const (
	Serialize_JSON = "json"
)

type DBGetterFunc func() interface{}

type SimpleCache struct {
	Operator  *StringOperator
	Expire    time.Duration
	DBGetter  DBGetterFunc
	Serialize string
}

func NewSimpleCache(operator *StringOperator, expire time.Duration, serialize string) *SimpleCache {
	return &SimpleCache{Operator: operator, Expire: expire, Serialize: serialize}
}



func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operator.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

func (this *SimpleCache) GetCache(key string) interface{} {
	ret := ""
	if this.Serialize == Serialize_JSON {
		f := func() string {
			obj := this.DBGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = this.Operator.Get(key).UnwrapOrElse(f)
		this.SetCache(key, ret)
	}
	return ret
}
