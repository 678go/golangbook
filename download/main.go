package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type FileDownloader struct {
	fileSize   int
	dir        string
	fileName   string
	url        string
	md5        string
	blockSize  int
	retryTimes int
}

func main() {
	now := time.Now()
	f := FileDownloader{
		url:        "https://go.dev/dl/go1.19.9.windows-amd64.zip",
		fileName:   "go1.19.9.windows-amd64.zip",
		dir:        "/Users/ylinyang/Downloads/test/2023/",
		blockSize:  5 * 1024 * 1024,
		retryTimes: 3,
		md5:        "3b0ca22aedf5fd85e84c944dd96ab3044213bfd224cc3e9850ad86f1f71e1be93",
	}
	req, err := http.NewRequest("GET", f.url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	var res *http.Response
	for i := 0; i < f.retryTimes; i++ {
		res, err = http.DefaultClient.Do(req)
		if err == nil {
			break
		}
		if i == f.retryTimes {
			panic(err)
		}
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(err)
	}
	fmt.Println(res)
	f.fileSize, err = strconv.Atoi(res.Header.Get("Content-Length"))
	fmt.Println("size", f.fileSize)

	// 创建本地临时文件
	tmpFile, err := os.CreateTemp(f.dir, "download_*.tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(f.dir + tmpFile.Name())
	defer tmpFile.Close()
	hasher := sha256.New()

	var downloadedSize int
	var buf = make([]byte, 1024)
	for {
		// 定义循环规范
		start, end := downloadedSize, downloadedSize+f.blockSize-1
		if end >= f.fileSize {
			end = f.fileSize - 1
		}
		// 根据已下载的数据位置发送 Range 请求
		fmt.Println("循环:", start, end)
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 写入数据到临时文件
		n, err := io.CopyBuffer(io.MultiWriter(tmpFile, hasher), res.Body, buf)
		res.Body.Close()
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println("缓冲区文件大小", n)

		downloadedSize += int(n)
		fmt.Printf("Downloaded: %d / %d bytes\n", downloadedSize, f.fileSize)

		// 下载完成后退出循环
		if downloadedSize == f.fileSize {
			break
		}
	}

	// 将下载好的临时文件改名为目标文件名
	err = os.Rename(tmpFile.Name(), f.dir+f.fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// hash验证
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	fmt.Println(f.md5)
	fmt.Println(time.Since(now))
}
