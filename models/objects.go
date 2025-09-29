package models

import "github.com/gofrs/uuid"

type Player struct {
	ID       uuid.UUID   `json:"id"`
	Login    string      `json:"login"`
	Password string      `json:"-"`
	Nickname string      `json:"nickname"`
	Stats    PlayerStats `json:"stats"`
}

type PlayerStats struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Draws  int `json:"draws"`
}

type RoomPhase int

const (
	RoomPhaseOpen RoomPhase = 0
	RoomPhaseFull RoomPhase = 1
)

type RoomPlayer struct {
	ID             uuid.UUID `json:"id"`
	Nickname       string    `json:"nickname"`
	RequestNewGame bool      `json:"-"`
}

type Room struct {
	ID          uuid.UUID   `json:"id"`
	Host        RoomPlayer  `json:"host"`
	Guest       *RoomPlayer `json:"guest"`
	GameID      *uuid.UUID  `json:"gameId"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Phase       RoomPhase   `json:"phase"`
}

type GamePhase int

const (
	GamePhaseInProgress GamePhase = 0
	GamePhaseCompleted  GamePhase = 1
)

type GamePlayer struct {
	ID   uuid.UUID `json:"id"`
	Mark string    `json:"mark"`
}

type Game struct {
	ID              uuid.UUID  `json:"id"`
	Host            GamePlayer `json:"host"`
	Guest           GamePlayer `json:"guest"`
	CurrentPlayerID uuid.UUID  `json:"currentPlayerId"`
	Board           string     `json:"board"`
	WinnerID        *uuid.UUID `json:"winnerId"`
	Phase           GamePhase  `json:"phase"`
}
