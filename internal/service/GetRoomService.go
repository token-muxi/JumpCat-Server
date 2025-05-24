package service

import (
	"database/sql"
	"errors"
)

type GetRoomService struct {
	Database *sql.DB
}

type Room struct {
	Room    int       `json:"room"`
	P1      string    `json:"p1"`
	P2      string    `json:"p2"`
	IsStart bool      `json:"is_start"`
	Map     *struct{} `json:"map"`
}

func NewGetRoomService(db *sql.DB) *GetRoomService {
	return &GetRoomService{Database: db}
}

func (s *GetRoomService) GetRoom(room int) (Room, error) {
	query := `SELECT p1, p2, is_start, map FROM room WHERE room = ?`
	row := s.Database.QueryRow(query, room)

	var p1, p2 string
	var isStart bool
	var mapData *struct{}
	err := row.Scan(&p1, &p2, &isStart, &mapData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Room{}, nil
		}
		return Room{}, err
	}

	if p2 == "" {
		mapData = nil
	} else {
		mapData = &struct{}{}
	}

	roomData := Room{
		Room:    room,
		P1:      p1,
		P2:      p2,
		IsStart: isStart,
		Map:     mapData,
	}

	return roomData, nil
}
