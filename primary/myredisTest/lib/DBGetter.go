package lib

import (
	"jtthinkStudy/myredis/models"
	"log"
)

func DBGetter(id string) func() interface{} {
	return func() interface {}{
		log.Println("get from db")
		item := models.NewItemModel()
		DB.Where("id = ?", id).First(item)
		//b, _ := json.Marshal(item)
		return item
	}
}
