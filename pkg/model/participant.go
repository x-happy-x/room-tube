package model

import "time"

type RoomParticipant struct {
	RoomID    int64     `json:"room_id" db:"room_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Role      string    `json:"role" db:"role"`
	Status    string    `json:"status" db:"status"`
	Deleted   bool      `json:"deleted" db:"deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
