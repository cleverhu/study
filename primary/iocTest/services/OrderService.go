package services

import "fmt"

type OrderService struct {
	DBService *DBService `inject:"-"`
	Version   string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "v1.0"}
}

func (this *OrderService) GetOrderInfo(orderID int) {
	fmt.Printf("获取用户订单信息:%d\n", orderID)
}
