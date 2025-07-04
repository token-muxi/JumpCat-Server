package service

import (
	"JumpCat-Server/middleware"
	"database/sql"
	"fmt"
)

type SubmitPrepareService struct {
	Database *sql.DB
}

func NewSubmitPrepareService(db *sql.DB) *SubmitPrepareService {
	return &SubmitPrepareService{Database: db}
}

func (s *SubmitPrepareService) UpdateStatus(room int, role string, Status bool) error {
	query := "UPDATE room SET " + role + "_ready = ? WHERE room = ?"
	_, err := s.Database.Exec(query, Status, room)
	if err != nil {
		middleware.Logger.Log("ERROR", fmt.Sprintf("failed to update status of room %d: %s", room, err.Error()))
		return err
	}
	return nil
}
