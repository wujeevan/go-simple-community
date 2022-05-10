package repository

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPostDao_QueryPostById(t *testing.T) {
	filePath := "../data/"
	if err := Init(filePath); err != nil {
		fmt.Println(err)
	}
	post := &Post{ParentID: 5, Content: "发布帖子"}
	// NewPostDaoInstance().InsertPostByParentId(filePath, post.ParentID, post)
	for i := 0; i < 10; i++ {
		go NewPostDaoInstance().QueryPostByParentId(int64(rand.Intn(2) + 1))
		go NewPostDaoInstance().InsertPostByParentId(filePath, post.ParentID, post)
	}
	time.Sleep(3 * time.Second)
}
