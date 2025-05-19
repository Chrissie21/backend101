package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"-"`
	Amount      float64   `json:"amount" validate:"required,gt=0"`
	Category    string    `json:"category" validate:"required,min=2,max=30"`
	Description string    `json:"description" validate:"required,min=2"`
	Type        string    `json:"type" validate:"required,oneof=income expense"` // income or expense
	Date        time.Time `json:"date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
