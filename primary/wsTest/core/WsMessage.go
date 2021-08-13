package core

import (
	"encoding/json"
	"jtthinkStudy/wsTest/models"
)

type WsMessage struct {
	messageType int
	data        []byte
}

func NewWsMessage(messageType int, data []byte) *WsMessage {
	return &WsMessage{messageType: messageType, data: data}
}

func (this *WsMessage) parseForCmd() (*models.WsResponse, error) {
	cmd := NewWsCommand()
	err := json.Unmarshal(this.data, cmd)
	if err != nil {
		return nil, err
	}
	return cmd.Parse()
}
