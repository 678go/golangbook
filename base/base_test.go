package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	//rand.Seed(1)
	fmt.Println(rand.Int())
}
func TestB(tt *testing.T) {
	TIME_FMT := "2006-01-02 15:04:05"
	now := time.Now()
	ts := now.Format(TIME_FMT) // 2023-03-12 16:00:53
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TIME_FMT, ts, loc)
	fmt.Println(ts, loc, t)
}
