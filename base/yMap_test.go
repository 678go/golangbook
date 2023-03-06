package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMap(t *testing.T) {
	arr := make([]int, 0, 10)
	// 空的struct是不占用内存的
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	for _, i := range arr {
		m[i] = struct{}{}
	}
	fmt.Println(len(m))
}
