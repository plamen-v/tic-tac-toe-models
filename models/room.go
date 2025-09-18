package models

type RoomPhase int

const (
	RoomPhaseOpen RoomPhase = 0
	RoomPhaseFull RoomPhase = 1
)

type RoomParticipant struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Continue bool   `json:"-"`
}

type Room struct {
	ID          int64            `json:"id"`
	Host        RoomParticipant  `json:"host"`
	Guest       *RoomParticipant `json:"guest"`
	GameID      *int64           `json:"gameId"`
	PrevGameID  *int64           `json:"prevGameId"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Phase       RoomPhase        `json:"phase"`
}
