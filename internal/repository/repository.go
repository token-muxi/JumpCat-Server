package repository

import (
	"JumpCat-Server/middleware"
	"database/sql"
	"encoding/json"
	"fmt"
)

type RoomRepository struct {
	db *sql.DB
}
type Room struct {
	RoomID    int             `json:"room"`
	Player1   string          `json:"p1"`
	Player2   string          `json:"p2"`
	Map       json.RawMessage `json:"map"`
	IsStarted bool            `json:"is_started"`
}

func NewRoomRepository(db *sql.DB)*RoomRepository{
	return &RoomRepository{
		db:db,
	}
}
// 创建房间
func (r *RoomRepository) InsertRoom(RoomMessage *Room) error {
	query := "INSERT INTO room (room, p1,map,is_started) VALUES (?,?,?,?)"
	_, err := r.db.Exec(query, RoomMessage.RoomID, RoomMessage.Player1, RoomMessage.Player2, RoomMessage.Map, RoomMessage.IsStarted)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to create room message:%s", err.Error()))
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

// 插入第二个用户
func (r *RoomRepository) InsertPlayer2(P2 string, RoomID int) error {
	query := "UPDATE room SET p2 = ? WHERE room = ?"
	_, err := r.db.Exec(query, P2, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to insert player2 of room:%d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}

// 更新状态
func (r *RoomRepository) UpdateStatus(RoomID int, Status bool) error {
	query := "UPDATE room SET is_started = ? WHERE room = ?"
	_, err := r.db.Exec(query, Status, RoomID)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to update status of room:%d:%s", RoomID, err.Error()))
		return err
	}
	return nil
}
