package repository

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"encoding/json"
)

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao   *TopicDao
	topicOnce  sync.Once
	topicMutex sync.RWMutex
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	topicMutex.RLock()
	defer topicMutex.RUnlock()
	fmt.Println(id, *topicIndexMap[id])
	return topicIndexMap[id]
}

func (*TopicDao) InsertNewTopic(filePath string, topic *Topic) error {
	topicMutex.Lock()
	defer topicMutex.Unlock()
	topic.ID = getNextTopicId()
	topic.CreateTime = time.Now().Unix()
	topicIndexMap[topic.ID] = topic
	if err := WriteTopic(filePath, topic); err != nil {
		return err
	}
	return nil
}

func getNextTopicId() int64 {
	nextTopicId := len(topicIndexMap) + 1
	return int64(nextTopicId)
}

func WriteTopic(filePath string, topic *Topic) error {
	open, err := os.OpenFile(filePath+"topic", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer open.Close()
	write := bufio.NewWriter(open)
	buf, err := json.Marshal(topic)
	if err != nil {
		return err
	}
	data := bytes.NewBuffer(buf).String()
	write.WriteString(data + "\n")
	if err := write.Flush(); err != nil {
		return err
	}
	return nil
}
