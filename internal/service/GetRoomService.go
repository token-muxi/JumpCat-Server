package service

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type GetRoomService struct {
	Database *sql.DB
}

func NewGetRoomService(db *sql.DB) *GetRoomService {
	return &GetRoomService{Database: db}
}

func (s *GetRoomService) GetRoom(room int) (Room, error) {
	query := `SELECT p1, p2, is_start, map FROM room WHERE room = ?`
	row := s.Database.QueryRow(query, room)

	var p1, p2 string
	var isStart bool
	var mapJSON string
	err := row.Scan(&p1, &p2, &isStart, &mapJSON)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Room{}, nil
		}
		return Room{}, err
	}

	var locData []Location
	if mapJSON != "" {
		locData = []Location{}
		if err := json.Unmarshal([]byte(mapJSON), &locData); err != nil {
			return Room{}, err
		}
	}

	roomData := Room{
		Room:    room,
		P1:      p1,
		P2:      p2,
		IsStart: isStart,
		Map:     locData,
	}

	return roomData, nil
}
