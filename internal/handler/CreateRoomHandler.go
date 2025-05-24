package handler

import (
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/repository"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/internal/util"
	"encoding/json"
	"net/http"
)

type CreateRoomRequest struct {
	Palyer1UUid string `json:"uuid"`
}

type CreateRoomReponse struct {
	RoomID int `json:"room_id"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	rr := repository.NewRoomRepository(db)
	s := service.NewCreateRoomService(rr)
	var request CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "invalid request")
		return
	}
	Room, err := s.CreateRoom(request.Palyer1UUid)
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "failed to create room")
		return
	}
	util.WriteResponse(w, http.StatusOK, CreateRoomReponse{
		RoomID: Room,
	})
}
