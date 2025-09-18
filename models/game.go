package models

type GamePhase int

const (
	GamePhaseStarted GamePhase = 0
	GamePhaseDone    GamePhase = 1
)

type GamePlayer struct {
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
	LoserID         *int64
	Phase           GamePhase
}
