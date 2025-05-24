package handler

import (
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/internal/util"
	"JumpCat-Server/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func GetRoom(w http.ResponseWriter, r *http.Request) {
	room, _ := strconv.Atoi(r.URL.Query().Get("room"))
	if room == 0 {
		util.WriteResponse(w, http.StatusBadRequest, "room is required")
		return
	}

	db := database.GetDB()
	if db == nil {
		middleware.Logger.Log("ERROR", "Database not initialized")
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// 创建实例
	getRoomService := service.NewGetRoomService(db)

	// 读取房间数据
	roomData, err := getRoomService.GetRoom(room)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to get room data: %s", err))
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	if roomData.Room == 0 {
		util.WriteResponse(w, http.StatusNotFound, "Room not found")
		return
	}

	util.WriteResponse(w, http.StatusOK, roomData)
}
