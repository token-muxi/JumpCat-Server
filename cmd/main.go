package main

import (
	"JumpCat-Server/internal/config"
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/router"
	"JumpCat-Server/middleware"
	"fmt"
	"net/http"
)

func main() {
	var err error

	// 加载配置
	cfg := config.LoadConfig()
	middleware.Logger = middleware.NewLogger(cfg)

	// 初始化路由
	r := router.NewRouter()
	loggedRouter := middleware.Logger.HttpMiddleware(r)

	// 初始化数据库连接
	err = database.NewDB(cfg)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to initialize database: %s", err))
		return
	}
	defer database.GetDB().Close()

	// 初始化 Redis 连接
	err = database.NewRedis(cfg)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to initialize Redis: %s", err))
		return
	}

	// 启动服务
	middleware.Logger.Log("INFO", fmt.Sprintf("Starting server on port %s", cfg.Port))
	err = http.ListenAndServe(":"+cfg.Port, loggedRouter)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to start server: %s", err))
	}
}
