package dto

type CreateTransactionInput struct {
	Amount      float64 `json:"amount" binding:"required"`
	Type        string  `json:"type" binding:"required,oneof=income expense"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description"`
}

type UpdateTransactionInput struct {
	Amount      float64 `json:"amount" binding:"required"`
	Type        string  `json:"type" binding:"required,oneof=income expense"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description"`
	// Date        string  `json:"date"` // Optional if allowing custom dates
}
