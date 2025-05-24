package service

import (
	"JumpCat-Server/middleware"
	"database/sql"
	"fmt"
)

type JoinRoomService struct {
	Database *sql.DB
}

func NewJoinService(db *sql.DB) *JoinRoomService {
	return &JoinRoomService{Database: db}
}

func (s *JoinRoomService) InsertPlayer2(P2 string, RoomID int) error {
	query := "UPDATE room SET p2 = ? WHERE room = ?"
	_, err := s.Database.Exec(query, P2, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert player2 of room:%d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}
