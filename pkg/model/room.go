package model

import "time"

type Room struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	IsPublic  bool      `json:"is_public" db:"is_public"`
	Password  string    `json:"password" db:"password"`
	Deleted   bool      `json:"deleted" db:"deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
