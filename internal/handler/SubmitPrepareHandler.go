package handler

import (
	"JumpCat-Server/internal/database"
	"JumpCat-Server/internal/service"
	"JumpCat-Server/internal/util"
	"JumpCat-Server/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

func SubmitPrepare(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Role string `json:"role"`
		Room int    `json:"room"`
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

	// 创建实例
	SubmitPrepareService := service.NewSubmitPrepareService(db)

	// 更新状态
	err = SubmitPrepareService.UpdateStatus(requestData.Room, requestData.Role, true)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to submit player %s status: %s", requestData.Role, err))
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}
	util.WriteResponse(w, http.StatusOK, fmt.Sprintf("update player %s success", requestData.Role))
}
