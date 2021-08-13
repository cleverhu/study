package main

import (
	"fmt"
	"jtthinkStudy/iocTest/injector"
	"jtthinkStudy/iocTest/services"
)

func main() {
	beanFactory := injector.BeanFactory
	beanFactory.Set(services.NewOrderService(), services.NewDBService())

	userService := services.NewUserService()
	//fmt.Println(userService)
	beanFactory.Apply(userService)
	fmt.Println(userService.Order.DBService)
}
