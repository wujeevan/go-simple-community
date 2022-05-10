package service

import (
	"errors"

	"github.com/wujeevan/go-simple-community/repository"
)

type TopicInfo = *repository.Topic

type InsertTopicFlow struct {
	Title     string
	Content   string
	TopicInfo *repository.Topic
}

func InsertTopic(title, content string) (TopicInfo, error) {
	return NewInsertTopicFlow(title, content).Do()
}

func NewInsertTopicFlow(title, content string) *InsertTopicFlow {
	return &InsertTopicFlow{
		Title:   title,
		Content: content,
	}
}

func (f *InsertTopicFlow) Do() (TopicInfo, error) {
	if err := f.CheckParam(); err != nil {
		return nil, err
	}
	if err := f.PrepareInfo(); err != nil {
		return nil, err
	}
	if err := f.PackPageInfo(); err != nil {
		return nil, err
	}
	return f.TopicInfo, nil
}

func (f *InsertTopicFlow) CheckParam() error {
	if len(f.Content) == 0 || len(f.Title) == 0 {
		return errors.New("topic's title or content cannot be empty")
	}
	return nil
}

func (f *InsertTopicFlow) PrepareInfo() error {
	TopicInfo, err := repository.NewTopicDaoInstance().InsertNewTopic("./data/", f.Title, f.Content)
	if err != nil {
		return err
	}
	f.TopicInfo = TopicInfo
	return nil
}

func (f *InsertTopicFlow) PackPageInfo() error {
	return nil
}
