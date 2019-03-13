package models

import (
	"time"
)

type User struct {
	ID uint `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
