package models

type Player struct {
	ID       int64       `json:"id"`
	Login    string      `json:"login"`
	Password string      `json:"-"`
	Nickname string      `json:"nickname"`
	Stats    PlayerStats `json:"stats"`
	RoomID   *int64      `json:"roomId"`
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

type RoomParticipant struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Continue bool   `json:"-"`
}

type Room struct {
	Host        RoomParticipant  `json:"host"`
	Guest       *RoomParticipant `json:"guest"`
	GameID      *int64           `json:"gameId"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Phase       RoomPhase        `json:"phase"`
}

type GamePhase int

const (
	GamePhaseInProgress GamePhase = 0
	GamePhaseCompleted  GamePhase = 1
)

type GamePlayer struct { //todo!
	ID   int64
	Mark rune
}

type Game struct {
	ID              int64
	HostID          int64
	HostMark        string
	GuestID         int64
	GuestMark       string
	CurrentPlayerID int64
	Board           string
	WinnerID        *int64
	Phase           GamePhase
}
