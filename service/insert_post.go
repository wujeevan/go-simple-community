package service

import (
	"errors"

	"github.com/wujeevan/go-simple-community/repository"
)

type PostInfo = *repository.Post

type InsertPostFlow struct {
	TopicId  int64
	Content  string
	PostInfo *repository.Post
}

func InsertPost(topicId int64, content string) (PostInfo, error) {
	return NewInsertPostFlow(topicId, content).Do()
}

func NewInsertPostFlow(topicId int64, content string) *InsertPostFlow {
	return &InsertPostFlow{
		TopicId: topicId,
		Content: content,
	}
}

func (f *InsertPostFlow) Do() (PostInfo, error) {
	if err := f.CheckParam(); err != nil {
		return nil, err
	}
	if err := f.PrepareInfo(); err != nil {
		return nil, err
	}
	if err := f.PackPageInfo(); err != nil {
		return nil, err
	}
	return f.PostInfo, nil
}

func (f *InsertPostFlow) CheckParam() error {
	if f.TopicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *InsertPostFlow) PrepareInfo() error {
	postInfo, err := repository.NewPostDaoInstance().InsertPostByParentId("./data/", f.TopicId, f.Content)
	if err != nil {
		return err
	}
	f.PostInfo = postInfo
	return nil
}

func (f *InsertPostFlow) PackPageInfo() error {
	return nil
}
