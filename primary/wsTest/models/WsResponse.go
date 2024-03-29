package models

import (
	"encoding/json"
)

type WsResponse struct {
	Type   string
	Result interface{}
}

func NewWsResponse(t string, result interface{}) *WsResponse {
	return &WsResponse{Type: t, Result: result}
}

func (this *WsResponse) ToJson() []byte {
	data, err := json.Marshal(this)
	if err != nil {
		return []byte("")
	}

	return data
}
