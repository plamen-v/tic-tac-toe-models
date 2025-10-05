package models

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateRoomRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RoomListRequest struct {
	PageInfo PageInfo `json:"pageInfo"`
}

type RankingRequest struct {
	PageInfo PageInfo `json:"pageInfo"`
}
