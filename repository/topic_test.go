package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTopicDao_QueryTopicById(t *testing.T) {
	filePath := "../data/"
	if err := Init(filePath); err != nil {
		fmt.Println(err)
	}
	topic := &Topic{Title: "Test", Content: "context"}
	fmt.Println(*topic)
	topic_json, err := json.Marshal(topic)
	if err != nil {
		fmt.Println(err)
	}
	data := bytes.NewBuffer(topic_json).String()
	fmt.Println(data)
	for i := 0; i < 10; i++ {
		go NewTopicDaoInstance().QueryTopicById(int64(rand.Intn(2) + 1))
		go NewTopicDaoInstance().InsertNewTopic(filePath, topic)
	}
	time.Sleep(time.Second)
}
