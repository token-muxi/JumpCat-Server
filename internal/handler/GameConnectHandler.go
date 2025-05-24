package handler

import (
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/middleware"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GameConnectHandler(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	roomID := r.URL.Query().Get("room")
	playerID := r.URL.Query().Get("player")
	if roomID == "" || playerID == "" {
		middleware.Logger.Log("ERROR", "Missing parameter")
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}

	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		middleware.Logger.Log("ERROR", "Invalid room ID")
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// 建立Redis连接
	redisClient := database.GetRedis()
	if redisClient == nil {
		middleware.Logger.Log("ERROR", "Redis client not initialized")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 建立数据库连接
	db := database.GetDB()
	if db == nil {
		middleware.Logger.Log("ERROR", "Database not initialized")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 获取玩家信息
	roomService := service.NewGetRoomService(db)
	roomData, err := roomService.GetRoom(roomIDInt)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to get room data: %s", err))
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// 确定对方ID
	var otherPlayerID string
	if roomData.P1 == playerID {
		otherPlayerID = roomData.P2
	} else if roomData.P2 == playerID {
		otherPlayerID = roomData.P1
	}

	// 建立WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("WebSocket upgrade failed: %s", err))
		return
	}
	defer conn.Close()

	middleware.Logger.Log("INFO", fmt.Sprintf("WebSocket connection established for room: %s, player: %s", roomID, playerID))
	ctx := context.Background()

	// 生成Redis键名
	playerKey := fmt.Sprintf("room:%s:player:%s", roomID, playerID)
	otherPlayerKey := fmt.Sprintf("room:%s:player:%s", roomID, otherPlayerID)

	// WebSocket消息循环
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			middleware.Logger.Log("ERROR", fmt.Sprintf("Error reading message: %s", err))
			// 连接关闭时，清除Redis中的位置数据
			redisClient.Del(ctx, playerKey)
			break
		}

		middleware.Logger.Log("DEBUG", fmt.Sprintf("Received message: %s", message))

		// 存储当前玩家位置到Redis
		err = redisClient.Set(ctx, playerKey, message, 0).Err()
		if err != nil {
			middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to save player location: %s", err))
			continue
		}

		// 获取对方玩家数据
		otherMessage, err := redisClient.Get(ctx, otherPlayerKey).Result()
		if err != nil {
			middleware.Logger.Log("INFO", "Other player location not found")
			continue
		}

		// 发送对方玩家数据
		if err = conn.WriteMessage(websocket.TextMessage, []byte(otherMessage)); err != nil {
			middleware.Logger.Log("ERROR", fmt.Sprintf("Error sending message: %s", err))
			break
		}
		middleware.Logger.Log("INFO", fmt.Sprintf("Sent other player data: %s to player: %s", otherMessage, playerID))
	}
}
