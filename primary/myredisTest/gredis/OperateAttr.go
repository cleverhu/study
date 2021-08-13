package gredis

import (
	"fmt"
	"time"
)

const (
	ATTR_EXPIRE = "expire" //过期时间
	ATTR_NX     = "nx"     //不存在
	ATTR_XX     = "xx"
)

type empty struct{}

type OperatorAttr struct {
	Name  string
	Value interface{}
}

type OperatorAttrs []*OperatorAttr

func WithExpire(t time.Duration) *OperatorAttr {
	return &OperatorAttr{
		Name:  ATTR_EXPIRE,
		Value: t,
	}
}

func WithNX() *OperatorAttr {
	return &OperatorAttr{
		Name:  ATTR_NX,
		Value: empty{},
	}
}

func WithXX() *OperatorAttr {
	return &OperatorAttr{
		Name:  ATTR_XX,
		Value: empty{},
	}
}

func (this OperatorAttrs) Find(name string) *InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}

	return NewInterfaceResult(nil, fmt.Errorf("find attrs error,attr:%s", name))
}
