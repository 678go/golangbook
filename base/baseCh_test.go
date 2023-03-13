package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func f(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func TestC(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) // 关闭ch 不能放，但能取
	f(ch)
}

func TestD(t *testing.T) {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				time.Sleep(2 * time.Second)
				fmt.Println(task)
			}
			//for i := range ch {
			//	fmt.Println(i)
			//}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 统计书的字数
func TestCountNumber(t *testing.T) {
	totalNum := 0
	totalWorkers := 100

	wg := sync.WaitGroup{}
	wg.Add(totalWorkers)
	for i := 0; i < totalWorkers; i++ {
		go func() {
			defer wg.Done()
			totalNum += 100
		}()
	}
	wg.Wait()
	fmt.Println("总字数：", totalNum)
}

func TestServer(t *testing.T) {

}
