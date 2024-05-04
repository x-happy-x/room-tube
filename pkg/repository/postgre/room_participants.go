package postgre

import (
	"database/sql"
	"errors"
	"tube/pkg/model"
)

func (r *Repository) GetParticipantByID(roomID, userID int64) (*model.RoomParticipant, error) {
	participant := &model.RoomParticipant{}
	err := r.DB.QueryRow("SELECT room_id, user_id, role, status, deleted, created_at, updated_at FROM room_participants WHERE room_id = ? AND user_id = ?", roomID, userID).Scan(&participant.RoomID, &participant.UserID, &participant.Role, &participant.Status, &participant.Deleted, &participant.CreatedAt, &participant.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("participant not found")
		}
		return nil, err
	}
	return participant, nil
}

func (r *Repository) AddParticipant(participant *model.RoomParticipant) error {
	_, err := r.DB.Exec("INSERT INTO room_participants (room_id, user_id, role, status, deleted, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", participant.RoomID, participant.UserID, participant.Role, participant.Status, participant.Deleted, participant.CreatedAt, participant.UpdatedAt)
	return err
}

func (r *Repository) UpdateParticipant(roomID, userID int64, status string) error {
	_, err := r.DB.Exec("UPDATE room_participants SET status = ?, updated_at = CURRENT_TIMESTAMP WHERE room_id = ? AND user_id = ?", status, roomID, userID)
	return err
}

func (r *Repository) RemoveParticipant(roomID, userID int64) error {
	_, err := r.DB.Exec("UPDATE room_participants SET deleted = true WHERE room_id = ? AND user_id = ?", roomID, userID)
	return err
}
