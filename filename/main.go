package main

import (
	"log"
	"os"
	"strings"
)

// 批量修改文件名
func main() {
	path := "/Users/ylinyang/Desktop/gin"
	changeFileName(path)
}
func changeFileName(path string) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		log.Println("打开目录失败", err)
	}
	for _, entry := range dirEntries {
		fileName := path + "/" + entry.Name()

		// 添加名字
		if err := os.Rename(fileName, path+"/"+strings.ReplaceAll(entry.Name(), ".mp4", "")+"(678go.xyz-专注精品资源的网站).mp4"); err != nil {
			log.Println("修改名字失败:", err)
			return
		}

		// 替换名字
		//if err := os.Rename(fileName, path+"/"+strings.ReplaceAll(entry.Name(), "((678go.xyz-专注精品资源的网站).mp4)", ".mp4")); err != nil {
		//	log.Println("替换失败: ", err)
		//	return
		//}
		// 是文件夹递归一下
		if entry.IsDir() {
			changeFileName(fileName)
		}
	}
}
