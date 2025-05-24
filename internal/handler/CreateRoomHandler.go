package handler

import (
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/internal/util"
	"JumpCat-Server/middleware"
	"encoding/json"
	"net/http"
)

type CreateRoomReponse struct {
	RoomID int `json:"room_id"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Palyer1UUid string `json:"uuid"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	db := database.GetDB()
	if db == nil {
		middleware.Logger.Log("ERROR", "Database not initialized")
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// 初始化实例
	rr := service.NewRoomRepository(db)
	s := service.NewCreateRoomService(rr)

	Room, err := s.CreateRoom(requestData.Palyer1UUid)
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "failed to create room")
		return
	}

	util.WriteResponse(w, http.StatusOK, CreateRoomReponse{
		RoomID: Room,
	})
}
