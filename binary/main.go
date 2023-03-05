package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	_
	g
)
const (
	h = iota
	i = 10
	j = iota
	k
)

type slice struct {
}

func main() {
	fmt.Println("5 二进制：", binaryFormat(5))
	fmt.Println("-5 二进制：", binaryFormat(-5))
	/*
		     5 二进制： 00000000000000000000000000000101
			-5 二进制： 11111111111111111111111111111011
	*/
	fmt.Println("-260 << 23 二进制：", binaryFormat(-260<<23))
	// 01111110000000000000000000000000 此时就是正数

	fmt.Printf("%f\n", math.Pi)   // 3.141593
	fmt.Printf("%.2f\n", math.Pi) // 3.14

}

// 输出一个int64对应的二进制
func binaryFormat(n int64) string {
	// 用于做字符串拼接的 效果比 + 好
	sb := strings.Builder{}
	// math.Pow(2, 31) 构建一个32位二进制 最高位为1 其余为0
	// 只想单纯保存为二进制 将其转为64位bit
	c := int64(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		// 判断n与当前c 做与运算 由于c只有一位是1，以此来判断当前是否是1
		if n&c != 0 {
			// sb字符串追加1
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 // 将1右移
	}
	return sb.String()
}
