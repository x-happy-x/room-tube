package repository

import "tube/pkg/model"

type Repository interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error

	GetRoomByID(id int64) (*model.Room, error)
	CreateRoom(room *model.Room) error
	UpdateRoom(room *model.Room) error
	DeleteRoom(id int64) error

	GetParticipantByID(roomID, userID int64) (*model.RoomParticipant, error)
	AddParticipant(participant *model.RoomParticipant) error
	UpdateParticipant(roomID, userID int64, status string) error
	RemoveParticipant(roomID, userID int64) error
}
