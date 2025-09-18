package requests

import "github.com/plamen-v/tic-tac-toe-models/models"

type GetOpenRoomsRequest struct {
	Host        string           `json:"host"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Phase       models.RoomPhase `json:"phase"`
}
