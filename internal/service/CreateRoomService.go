package service

import (
	"JumpCat-Server/middleware"
	"fmt"
	"math/rand"
	"time"
)

type CreateRoomService struct {
	RoomRepository *RoomRepository
}

type Map struct {
	Length int64      `json:"length"`
	Locas  []Location `json:"locations"`
}

type Location struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
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
	MapMessage := InitMap()

	RoomMessage := Room{
		Room:     RoomID,
		P1:       Player1,
		Map:      MapMessage,
		P1_ready: false,
		P2_ready: false,
	}

	err := cs.RoomRepository.InsertRoom(&RoomMessage)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert map message for room %d: %s", RoomID, err))
		return 0, err
	}
	return RoomID, nil
}

// InitMap 生成地图
// TODO: 添加生成限制逻辑
func InitMap() Map {
	length := rand.Intn(600) + 200 //地图长度200-600
	cat_length := 3
	current := 0
	count := rand.Intn(length/10) + 5 //地刺个数限制
	Locas := make([]Location, count)
	// 生成任意的地刺
	for i := 0; i < len(Locas); i++ {
		//地刺长度
		spike_length := rand.Intn(4) + 1
		Locas[i].Start = int64(rand.Intn(10) + current + cat_length)
		Locas[i].End = Locas[i].Start + int64(spike_length)
		current = int(Locas[i].End)
	}
	return Map{
		Length: int64(length),
		Locas:  Locas,
	}
}
