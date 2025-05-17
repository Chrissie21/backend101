package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"-"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // income or expense
	Date        time.Time `json:"date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
