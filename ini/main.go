package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	cfg, err := ini.Load("./ini/a.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
	// ...
	p := new(Person)
	err = cfg.MapTo(p)
	fmt.Println(p)
	// ...

	// 一切竟可以如此的简单。
	err = ini.MapTo(p, "a.ini")
	// ...

	// 嗯哼？只需要映射一个分区吗？
	n := new(Note)
	err = cfg.Section("Note").MapTo(n)
	// ...
}
