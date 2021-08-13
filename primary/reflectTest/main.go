package main

import (
	"errors"
	"fmt"
	"reflect"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func map2Struct(m map[string]interface{}, data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() != reflect.Ptr {
		return errors.New("Data must be ptr! ")
	}
	t = t.Elem()
	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("Data must be struct! ")
	}

	for i := 0; i < t.NumField(); i++ {
		t1 := t.Field(i)
		v1 := v.Field(i)

		d, exists := m[t1.Tag.Get("json")]
		if exists && reflect.TypeOf(d) == v1.Type() {
			v1.Set(reflect.ValueOf(d))
		}
	}
	return nil
}

func main() {
	u:=&user{}
	m := map[string]interface{}{
		"id":   1,
		"name": "Test",
		"Age":  18,
	}

	err := map2Struct(m, u)
	fmt.Println(err)
	fmt.Println(u)
}
