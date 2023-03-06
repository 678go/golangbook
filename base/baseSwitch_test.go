package main

import (
	"fmt"
	"testing"
)

func breakDemo() {
BREAKDEMO:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO
			}
			fmt.Printf("%v-%v", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue
			}
			fmt.Printf("%v-%v", i, j)
		}
	}
}

func TestGoTo(t *testing.T) {
	switchType()
}

func switchType() {
	var n interface{} = 5
	switch v := n.(type) {
	case int:
		fmt.Printf("int %d\n", v)
	case float32:
		fmt.Printf("float32 %f\n", v)
	default:
		fmt.Println("不支持这种类型")
	}
}
