package main

import (
	"fmt"
	"testing"
)

// 接口是一组行为规范的集合
type Transporter interface { // 通常以er结尾
	// 接口里面只定义方法，不定义变量
	move(src string, dst string) (int, error)
	// 参数列表和返回值列表里面的变量名可以省略
	whistle(int) int
}

type Steamer interface {
	Transporter // 嵌套该接口，相当于Transporter是Steamer的字集
	displacement() int
}

// 定义结构体时无需显式声明他实现了什么接口
type Car struct {
	price int
}

// 只要结构体拥有接口里声明的所有方法，就称该结构体 实现了接口
func (c Car) move(src string, dst string) (int, error) {
	return c.price, nil
}

func (c Car) whistle(n int) int {
	return n
}
func foo(a Transporter) {
	a.whistle(100)
}

func TestName(t *testing.T) {
	var tr Transporter
	var c Car
	tr = c
	fmt.Println(tr.whistle(10))
	foo(c)
}

type payer interface {
	pay(int) error
}

type wc struct{}

func (w *wc) pay(int) error {
	fmt.Println("wc 10")
	return nil
}

type zf struct{}

func (z *zf) pay(int) error {
	fmt.Println("zf 20")
	return nil
}

func CheckOut(obj payer) {
	obj.pay(100)
}

func main1() {
	CheckOut(&wc{})
	CheckOut(&zf{})
}

func check(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println("i is int", v)
	} else {
		fmt.Println("i is not int", v)
	}
}
