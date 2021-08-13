package gredis

import (
	"context"
	"time"
)

type StringOperator struct {
	ctx context.Context
}

func NewStringOperator() *StringOperator {
	return &StringOperator{ctx: context.Background()}
}

func (this *StringOperator) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

func (this *StringOperator) MGet(keys ...string) *sliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, keys...).Result())
}

func (this *StringOperator) Set(key string, value interface{}, attrs ...*OperatorAttr) *InterfaceResult {
	exp := OperatorAttrs(attrs).Find(ATTR_EXPIRE)
	nx := OperatorAttrs(attrs).Find(ATTR_NX).UnwrapOr(nil)
	if nx != nil {
		return NewInterfaceResult(Redis().SetNX(this.ctx, key, value, exp.UnwrapOr(0*time.Second).(time.Duration)).Result())
	}
	xx := OperatorAttrs(attrs).Find(ATTR_XX).UnwrapOr(nil)
	if xx != nil {
		return NewInterfaceResult(Redis().SetXX(this.ctx, key, value, exp.UnwrapOr(0*time.Second).(time.Duration)).Result())
	}
	return NewInterfaceResult(Redis().Set(this.ctx, key, value, exp.UnwrapOr(0*time.Second).(time.Duration)).Result())
}
