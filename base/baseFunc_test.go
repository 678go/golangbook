package main

import "fmt"

// a,b是形参，函数内部的局部变量，实参的值会拷贝给形参
// 在形参类型相同时 可以只写一次 如：func add(a,b int) c{}
//func add(a *int, b *int) int {
//	*a = *a + *b
//	return *a
//}

type user struct {
	name  string
	hello func(name string) string
}

func main() {
	ch := make(chan func(name string) string, 0)

	ch <- func(name string) string {
		return "hello" + name
	}
}

func sub() func() {
	i := 10
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("%p\n", &i)
		i--
		fmt.Println(i)
	}
	return b
}

func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("%p\n", &base)
		base += i
		return base
	}
}

func deferExe() (i int) {
	i = 9
	defer func() { // 这里定义了一个匿名函数func(){} + () 这个扩号表示调用该函数 有参数则传参
		fmt.Printf("i=%d\n", i) // i = 5
	}()
	defer fmt.Printf("i=%d\n", i) // i = 9
	return 5                      // 当return 的时候会将 5 赋值给i
}
