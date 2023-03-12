package main

import (
	"fmt"
	"reflect"
	"testing"
)

type u struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func (s u) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s u) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func TestGetField(t *testing.T) {
	typeOf := reflect.TypeOf(u{})
	// 遍历机结构体所有字段
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		// name:name index:[0] type:string json tag:name
		// name:age  index:[1]  type:int    json tag:age
	}
	// 指定字段
	if fieldByName, ok := typeOf.FieldByName("age"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", fieldByName.Name, fieldByName.Index, fieldByName.Type, fieldByName.Tag.Get("json"))
		// name:age index:[1] type:int json tag:age
	}
	// 获取结构体的方法
	printMethod(u{})
	/*
	  method name:Sleep
	  method:func() string
	  好好睡觉，快快长大。
	  method name:Study
	  method:func() string
	  好好学习，天天向上。
	*/
}
