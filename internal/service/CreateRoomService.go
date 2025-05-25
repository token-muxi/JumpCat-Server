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
func InitMap() Map {
	length := 200 // 地图长度
	catLength := 3
	current := 0
	var Locas []Location

	// 直到剩余长度不足5
	for length-current >= 5 {
		// 地刺长度
		spikeLength := rand.Intn(4) + 1

		// 计算起始位置
		start := rand.Intn(3) + current + catLength
		end := start + spikeLength

		// 如果超出地图边界，停止生成
		if end >= length {
			break
		}

		// 添加新地刺
		Locas = append(Locas, Location{
			Start: int64(start),
			End:   int64(end),
		})

		current = end // 更新当前位置
	}

	return Map{
		Length: int64(length),
		Locas:  Locas,
	}
}
