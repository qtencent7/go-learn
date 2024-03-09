package test

import (
	"log"
	"net/http"
)

func Test() {
	// 使用 new 创建一个新的 http.Server 实例
	server := new(http.Server)

	// 设置服务器的监听地址和处理函数
	server.Addr = ":8080"
	server.Handler = http.NewServeMux()

	// 启动服务器
	log.Fatal(server.ListenAndServe())
}
