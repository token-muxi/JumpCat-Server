package service

import (
	"JumpCat-Server/internal/repository"
	"JumpCat-Server/middleware"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type CreateRoomService struct {
	RoomRepository *repository.RoomRepository
}

// 地图
type Map struct {
	Locas []Location `json:"map"`
}
type Location struct {
	X     int64 `json:"x"`
	Width int64 `json:"width"`
}

func NewCreateRoomService(r *repository.RoomRepository) *CreateRoomService {
	return &CreateRoomService{
		RoomRepository: r,
	}
}

func (cs *CreateRoomService) CreateRoom(Player1 string) (int, error) {
	//生成房间ID
	rand.Seed(time.Now().UnixNano())
	RoomID := rand.Intn(900000) + 100000
	Map := InitMap()
	MapMessage, err := json.Marshal(Map)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to serialize map message:%s", err))
		return 0, err
	}
	RoomMessage := repository.Room{
		RoomID:    RoomID,
		Player1:   Player1,
		Map:       MapMessage,
		IsStarted: false,
	}
	err = cs.RoomRepository.InsertRoom(&RoomMessage)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert room:%d:map message:%s", RoomID, err))
		return 0, err
	}
	return RoomID, nil
}

// 生成地图
// TODO:添加生成限制逻辑
func InitMap() Map {
	length := 100000
	current := 0
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(20) + 1
	//RoomID:=rand.Intn(900000)+100000
	Locas := make([]Location, count)
	//生成任意的地刺
	for i := 0; i < len(Locas); i++ {
		Locas[i].X = int64(rand.Intn(length-current+1) + current)
		current = int(Locas[i].X)
		Locas[i].Width = int64(rand.Intn(200) + 100) //生成范围100-300
	}
	return Map{
		Locas: Locas,
	}
}
