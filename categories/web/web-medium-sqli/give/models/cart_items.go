package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	UserID   uuid.UUID `json:"user_id"`
	ItemID   uint      `json:"item_id"`
	Name     string    `json:"name"`
	ImageURL string    `json:"image_url"`
	Quantity int       `json:"quantity"`
}
