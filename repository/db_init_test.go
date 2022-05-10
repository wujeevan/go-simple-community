package repository

import (
	"testing"
)

func Test_initTopicIndexMap(t *testing.T) {
	filePath := "../data/topic"
	initTopicIndexMap(filePath)
}

func Test_initPostIndexMap(t *testing.T) {
	filePath := "../data/post"
	initPostIndexMap(filePath)
}
