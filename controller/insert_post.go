package controller

import (
	"strconv"

	"github.com/wujeevan/go-simple-community/service"
)

type PostData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InsertPost(topicIdStr string, content string) *PostData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PostData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	post, err := service.InsertPost(topicId, content)
	if err != nil {
		return &PostData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PostData{
		Code: 0,
		Msg:  "success",
		Data: post,
	}
}
