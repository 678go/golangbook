package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestStringSubCount 子串的统计
func TestStringSubCount(t *testing.T) {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZWXYZ"
	fmt.Println(strings.Count(s, "WXYZ"))
}

// TestStringReplace 替换
func TestStringReplace(t *testing.T) {
	s := "ABCDEFGHIJKLM NOPQRSTUVWX YZWXYZ"
	replace := strings.Replace(s, " ", "%20", -1)
	fmt.Println(replace)
	// 入参：需要替换的字符串 替换前 替换后 替换的长度 负数表示不限制
}
