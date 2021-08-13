package core

import (
	"encoding/json"
	"fmt"
	"jtthinkStudy/wsTest/models"
	"reflect"
)

const (
	NewPod  = 101
	PodList = 102
)

var CommandModel = map[int]models.IModel{}

func init() {
	CommandModel = make(map[int]models.IModel)
	CommandModel[NewPod] = (*models.PodModel)(nil)
	CommandModel[PodList] = (*models.PodListModel)(nil)
}

type WsCommand struct {
	CmdType   int
	CmdData   map[string]interface{}
	CmdAction string
}

func NewWsCommand() *WsCommand {
	return &WsCommand{}
}

func (this *WsCommand) Parse() (*models.WsResponse, error) {
	if v, ok := CommandModel[this.CmdType]; ok {
		newObj := reflect.New(reflect.TypeOf(v).Elem()).Interface()
		b, _ := json.Marshal(this.CmdData)
		err := json.Unmarshal(b, newObj)
		if err != nil {
			return nil, err
		}
		return newObj.(models.IModel).ParseAction(this.CmdAction)
	}
	return nil, fmt.Errorf("error action")
}
