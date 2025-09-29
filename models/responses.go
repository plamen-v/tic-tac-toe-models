package models

import (
	"github.com/gofrs/uuid"
)

type ErrorResponse struct {
	Code    string `json:"errorCode"`
	Message string `json:"errorMessage"`
}

type LoginResponse struct {
	Player *Player `json:"player"`
}

type OpenRoomsResponse struct {
	Rooms []*Room `json:"rooms"`
}

type CreateRoomResponse struct {
	RoomID uuid.UUID `json:"roomId"`
}

type GameStateResponse struct {
	Game *Game `json:"game"`
}

type RankResponse struct {
	Players []*Player `json:"players"`
}
