package models

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"Somebody Someone"`
	Email    string `json:"email" binding:"required,email" example:"somebody@someone.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"somebody@someone.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}
