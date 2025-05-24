package service

import (
	"JumpCat-Server/middleware"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type CreateRoomService struct {
	RoomRepository *RoomRepository
}

type Location struct {
	X     int64 `json:"x"`
	Width int64 `json:"width"`
}

func NewCreateRoomService(r *RoomRepository) *CreateRoomService {
	return &CreateRoomService{
		RoomRepository: r,
	}
}

func (cs *CreateRoomService) CreateRoom(Player1 string) (int, error) {
	// 生成房间ID
	rand.Seed(time.Now().UnixNano())
	RoomID := rand.Intn(900000) + 100000
	MapMessage, err := json.Marshal(InitMap())
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to serialize map message: %s", err))
		return 0, err
	}

	var locData []Location
	if err := json.Unmarshal(MapMessage, &locData); err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to deserialize map message: %s", err))
		return 0, err
	}

	RoomMessage := Room{
		Room:    RoomID,
		P1:      Player1,
		Map:     locData,
		P1_ready: false,
		P2_ready: false,
	}

	err = cs.RoomRepository.InsertRoom(&RoomMessage)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert map message for room %d: %s", RoomID, err))
		return 0, err
	}
	return RoomID, nil
}

// InitMap 生成地图
// TODO: 添加生成限制逻辑
func InitMap() []Location {
	length := 100000
	current := 0
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(20) + 1
	Locas := make([]Location, count)
	// 生成任意的地刺
	for i := 0; i < len(Locas); i++ {
		Locas[i].X = int64(rand.Intn(length-current+1) + current)
		current = int(Locas[i].X)
		Locas[i].Width = int64(rand.Intn(200) + 100) // 生成范围100-300
	}
	return Locas
}
