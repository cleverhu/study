package main

import (
	"bytes"
	"fmt"
)

type empyt struct{}

type set map[interface{}]empyt

func (this set) Add(vs ...interface{}) set {
	fmt.Printf("%p\n", this)
	for _, v := range vs {
		this[v] = struct{}{}
	}
	return this
}

func (this set) String() string {
	b := bytes.Buffer{}
	for k := range this {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		b.WriteString(fmt.Sprintf("%v", k))
	}
	return b.String()
}

func (this set) delete(k interface{})  {
	delete(this, k)

}

func main() {
	s := set{}
	fmt.Printf("%p\n", s)

	s.Add(1, 2, 3, "a", "aa").delete("a")
	fmt.Println(s)
}
