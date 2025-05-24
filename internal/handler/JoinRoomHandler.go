package handler

import(
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/internal/util"
	"JumpCat-Server/middleware"
	"encoding/json"
	"net/http"
)

func JoinRoom(w http.ResponseWriter,r *http.Request){
	var requestData struct{
		UUid string `json:"uuid"`
		RoomID int `json:"room"`
	}

	err:=json.NewDecoder(r.Body).Decode(&requestData)
	if err!=nil{
		util.WriteResponse(w,http.StatusBadRequest,nil)
		return
	}

	db:=database.GetDB()
	if db == nil {
		middleware.Logger.Log("ERROR", "Database not initialized")
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	//初始化实例
	s:=service.NewJoinService(db)

	err=s.InsertPlayer2(requestData.UUid,requestData.RoomID)
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "failed to join the room")
		return
	}
	util.WriteResponse(w,http.StatusOK,"join room successfully!")

}