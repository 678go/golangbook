package main

import (
	"context"
	"testing"
)

func TestBucketStrategy(t *testing.T) {
	bucket := getFileType("B")
	bucketStrategy, _ := NewBucketStrategy(bucket)
	_ = bucketStrategy.upload(context.Background(), "localPath")
}

func getFileType(file string) string {
	if file == "A" {
		return "tencent"
	}
	return "ali"
}
