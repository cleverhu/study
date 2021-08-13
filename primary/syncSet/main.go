package main

import (
	"bytes"
	"fmt"
	"sync"
)

type empty struct{}

type set struct {
	data sync.Map
}

func (this *set) Add(vs ...interface{}) *set {
	fmt.Printf("%p\n",&this)
	for _, v := range vs {
		this.data.Store(v, empty{})
	}
	return this
}

func (this *set) String() string {
	b := bytes.Buffer{}
	this.data.Range(func(key, value interface{}) bool {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		b.WriteString(fmt.Sprintf("%v", key))
		return true
	})

	return b.String()

}

func (this *set) delete(key interface{}) *set {
	fmt.Printf("%p\n",this.data)
	this.data.Delete(key)
	return this
}

func main() {

	s := &set{}
	fmt.Printf("%p\n",s)
	s.Add(1, 2, 3, "a", "aa").delete(1)
	fmt.Printf("%p\n",&s)
	fmt.Println(s)

}
