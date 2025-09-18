package requests

import "github.com/plamen-v/tic-tac-toe/src/models"

type GetOpenRoomsRequest struct {
	Host        string           `json:"host"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Phase       models.RoomPhase `json:"phase"`
}
