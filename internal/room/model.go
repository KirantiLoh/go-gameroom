package room

import (
	"github.com/kirantiloh/gameroom/pkg/database"
)

type Room struct {
	RoomDto
	UUID      string            `json:"uuid"`
	ID        int               `json:"id"`
	CreatedAt database.JSONTime `json:"created_at"`
}

type RoomDto struct {
	CreateRoomDto
	LeaderID int `json:"leader_id"`
}

type CreateRoomDto struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}
