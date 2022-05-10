package controller

import (
	"github.com/wujeevan/go-simple-community/service"
)

type TopicData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InsertTopic(title, content string) *TopicData {
	topic, err := service.InsertTopic(title, content)
	if err != nil {
		return &TopicData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &TopicData{
		Code: 0,
		Msg:  "success",
		Data: topic,
	}
}
