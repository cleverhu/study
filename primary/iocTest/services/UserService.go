package services

import "fmt"

type UserService struct {
	Order *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserInfo(uid int) {
	fmt.Printf("获取用户信息:%d\n", uid)
}



func (this *UserService) GetOrderInfo(orderID int) {
	this.Order.GetOrderInfo(orderID)
}
