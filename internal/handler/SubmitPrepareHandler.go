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

func SubmitPrepare(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	room, _ := strconv.Atoi(r.URL.Query().Get("room"))
	if role == "" || room == 0 {
		util.WriteResponse(w, http.StatusBadRequest, "role and room is required")
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
	err := SubmitPrepareService.UpdateStatus(room, role, true)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("Failed to submit player %s status: %s", role, err))
		util.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}
	util.WriteResponse(w, http.StatusOK, fmt.Sprintf("update player %s success", role))
}
