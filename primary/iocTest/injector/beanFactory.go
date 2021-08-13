package injector

import (
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = &BeanFactoryImpl{beanMapper: make(BeanMapper)}
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
}

func (this *BeanFactoryImpl) Set(vs ...interface{}) {
	if vs == nil || len(vs) == 0 {
		return
	}
	for _, v := range vs {
		this.beanMapper.add(v)
	}
}

func (this *BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	getV := this.beanMapper.get(v)
	if getV.IsValid() {
		return getV.Interface()
	}
	return nil
}

func (this *BeanFactoryImpl) Apply(bean interface{}) {
	v := reflect.ValueOf(bean)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tField := v.Type().Field(i)

		if v.Field(i).CanSet() && tField.Tag.Get("inject") == "-" {
			getV := this.Get(tField.Type)
			if getV != nil {
				v.Field(i).Set(reflect.ValueOf(getV))
				this.Apply(getV)
			}

		}
	}

}
