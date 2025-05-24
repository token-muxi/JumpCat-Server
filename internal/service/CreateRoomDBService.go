package service

import (
	"JumpCat-Server/middleware"
	"database/sql"
	"encoding/json"
	"fmt"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

// InsertRoom 创建房间
func (r *RoomRepository) InsertRoom(RoomMessage *Room) error {
	query := "INSERT INTO room (room, p1, p2, is_start, map) VALUES (?, ?, ?, ?, ?)"

	jsonMap, err := json.Marshal(RoomMessage.Map)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to marshal map: %s", err.Error()))
		return err
	}

	_, err = r.db.Exec(query, RoomMessage.Room, RoomMessage.P1, RoomMessage.P2, RoomMessage.IsStart, string(jsonMap))
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to create room: %s", err.Error()))
		return err
	}
	return nil
}

func (r *RoomRepository) DeleteRoom(RoomID int) error {
	query := "DELETE FROM room WHERE room = ?"
	_, err := r.db.Exec(query, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to delete room:%d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}

// InsertPlayer2 插入第二个用户
func (r *RoomRepository) InsertPlayer2(P2 string, RoomID int) error {
	query := "UPDATE room SET p2 = ? WHERE room = ?"
	_, err := r.db.Exec(query, P2, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert player2 of room:%d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (r *RoomRepository) UpdateStatus(RoomID int, Status bool) error {
	query := "UPDATE room SET is_start = ? WHERE room = ?"
	_, err := r.db.Exec(query, Status, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to update status of room %d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}
