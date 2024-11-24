package models

import (
	"github.com/google/uuid"
)
type Profile struct {
	UserID  uuid.UUID
	Name    string
	Address string
	Phone   string
}
