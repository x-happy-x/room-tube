package postgre

import (
	"database/sql"
	"errors"
	"tube/pkg/model"
)

func (r *Repository) GetRoomByID(id int64) (*model.Room, error) {
	room := &model.Room{}
	err := r.DB.QueryRow("SELECT id, name, is_public, password, deleted, created_at, updated_at FROM rooms WHERE id = ?", id).Scan(&room.ID, &room.Name, &room.IsPublic, &room.Password, &room.Deleted, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return room, nil
}

func (r *Repository) CreateRoom(room *model.Room) error {
	_, err := r.DB.Exec("INSERT INTO rooms (name, is_public, password, deleted, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", room.Name, room.IsPublic, room.Password, room.Deleted, room.CreatedAt, room.UpdatedAt)
	return err
}

func (r *Repository) UpdateRoom(room *model.Room) error {
	_, err := r.DB.Exec("UPDATE rooms SET name = ?, is_public = ?, password = ?, deleted = ?, updated_at = ? WHERE id = ?", room.Name, room.IsPublic, room.Password, room.Deleted, room.UpdatedAt, room.ID)
	return err
}

func (r *Repository) DeleteRoom(id int64) error {
	_, err := r.DB.Exec("UPDATE rooms SET deleted = true WHERE id = ?", id)
	return err
}
