package pb

import (
	"fmt"
	"testing"
)
import "google.golang.org/protobuf/proto"

func TestP(t *testing.T) {
	// 序列化
	obj := &String{Value: "run go"}
	objBytes, _ := proto.Marshal(obj)
	fmt.Println(objBytes)

	// 反序列化
	uObj := &String{}
	_ = proto.Unmarshal(objBytes, uObj)
	fmt.Println(uObj.Value)
}
