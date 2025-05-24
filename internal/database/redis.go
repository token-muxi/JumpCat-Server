package database

import (
	"JumpCat-Server/internal/config"
	"JumpCat-Server/middleware"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// NewRedis 初始化 Redis 连接
func NewRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("[REDIS] Failed to connect: %s", err))
		return err
	}
	middleware.Logger.Log("INFO", "[REDIS] Connected successfully")
	return nil
}

// GetRedis 获取 Redis 客户端
func GetRedis() *redis.Client {
	return redisClient
}
