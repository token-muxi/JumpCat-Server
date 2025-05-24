package main

import (
	"JumpCat-Server/internal/router"
	"log"
	"net/http"
)

func main() {
	// 初始化路由
	r := router.NewRouter()

	// 启动服务
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
