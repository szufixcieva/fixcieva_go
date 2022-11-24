package main

import (
	"fmt"
	"net/http"
)

// 请求
func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		return
	}
}

func main() {
	// 定义请求路径及响应函数
	http.HandleFunc("/hello", hello)
	// 启动一个服务并监听 8080 端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
}
