package repository

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDao struct {
}

var (
	postDao   *PostDao
	postOnce  sync.Once
	postMutex sync.RWMutex
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostByParentId(parentId int64) []*Post {
	postMutex.RLock()
	defer postMutex.RUnlock()
	// fmt.Println(parentId, *postIndexMap[parentId][0])
	return postIndexMap[parentId]
}

func (*PostDao) InsertPostByParentId(filePath string, parentId int64, post *Post) error {
	postMutex.Lock()
	defer postMutex.Unlock()
	nextPostId, err := getNextPostId(parentId)
	if err != nil {
		return err
	}
	post.ID = nextPostId
	post.CreateTime = time.Now().Unix()
	if nextPostId == 1 {
		postIndexMap[parentId] = []*Post{post}
	} else {
		posts := postIndexMap[parentId]
		posts = append(posts, post)
		postIndexMap[parentId] = posts
	}
	WritePost(filePath, parentId, post)
	return nil
}

func getNextPostId(parentID int64) (int64, error) {
	if _, ok := topicIndexMap[parentID]; !ok {
		return 0, errors.New("must be posted in the existing topic")
	}
	posts, ok := postIndexMap[parentID]
	if !ok {
		return 1, nil
	}
	nextPostId := posts[len(posts)-1].ID + 1
	return nextPostId, nil
}

func WritePost(filePath string, paraentId int64, post *Post) error {
	open, err := os.OpenFile(filePath+"post", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer open.Close()
	write := bufio.NewWriter(open)
	buf, err := json.Marshal(post)
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
