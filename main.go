package main

import (
	"fmt"
	"go-learn/build_a_server"
	"net/http"
)

func test() {
	build_a_server.Router()
	// 使用corsMiddleware中间件包装默认的ServeMux
	http.Handle("/", build_a_server.Cors(build_a_server.R))

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
