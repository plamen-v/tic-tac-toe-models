package models

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateRoomRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetOpenRoomsRequest struct {
	Host        string    `json:"host"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Phase       RoomPhase `json:"phase"`
}
