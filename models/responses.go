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

type RoomResponse struct {
	Room *Room `json:"room"`
}

type RoomListResponse struct {
	Rooms    []*Room  `json:"rooms"`
	PageInfo PageInfo `json:"pageInfo"`
}

type CreateRoomResponse struct {
	RoomID uuid.UUID `json:"roomId"`
}

type GameResponse struct {
	Game *Game `json:"game"`
}

type RankingResponse struct {
	Players  []*Player `json:"players"`
	PageInfo PageInfo  `json:"pageInfo"`
}

type PageInfo struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	TotalCnt int `json:"totalCnt"`
}
