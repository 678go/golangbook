package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	fmt.Printf("s:%v len(s):%v cap(s):%v", s, len(s), cap(s))

}
