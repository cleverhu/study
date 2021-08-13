package services

import (
	"github.com/gin-gonic/gin"
	"jtthinkStudy/primary/eventBusTest/eventbus"
)

const GetProdList = "GetProdList"

func GetProdListCh() eventbus.EventDataChan {
	return Bus.Sub(GetProdList, NewProdService().GetProdList)
}

type ProdModel struct {
	ID   int
	Name string
}

type ProdService struct {
}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (this *ProdService) GetProdList(page int) gin.H {
	list := []ProdModel{{
		ID:   1,
		Name: "java",
	}, {
		ID:   2,
		Name: "golang",
	}}
	return gin.H{"data": gin.H{"result": list}, "page": page}
}
