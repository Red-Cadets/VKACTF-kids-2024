package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID           uint
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Comment      string    `json:"comment"`
	PromoCode    string    `json:"promo_code"`
	Items        []Cart    `json:"items"`
	IsDelivered  bool
	DeliveryTime int `json:"delivery_time"`
}
