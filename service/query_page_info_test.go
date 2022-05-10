package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wujeevan/go-simple-community/repository"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	topicId := 4
	pageInfo, err := QueryPageInfo(int64(topicId))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(*pageInfo.Topic)
	for _, post := range pageInfo.PostList {
		fmt.Println(*post)
	}
	assert.NotEqual(t, nil, pageInfo)
	assert.Equal(t, 0, len(pageInfo.PostList))
}
