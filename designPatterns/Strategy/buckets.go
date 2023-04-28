package main

import (
	"context"
	"fmt"
)

// BucketStrategy 上传策略
type BucketStrategy interface {
	upload(ctx context.Context, file string) error
}

// 定义bucket有哪些 进行选择
var buckets = map[string]BucketStrategy{
	"tencent": &tencentBucket{},
	"ali":     &aliBucket{},
}

func NewBucketStrategy(b string) (BucketStrategy, error) {
	s, ok := buckets[b]
	if !ok {
		return nil, fmt.Errorf("not found bucket %s", b)
	}
	return s, nil
}

// 腾讯云bucket
type tencentBucket struct{}

func (t *tencentBucket) upload(ctx context.Context, file string) error {
	fmt.Println("上传腾讯云")
	return nil
}

// 阿里云bucket
type aliBucket struct{}

func (a *aliBucket) upload(ctx context.Context, file string) error {
	fmt.Println("上传阿里云")
	return nil
}
