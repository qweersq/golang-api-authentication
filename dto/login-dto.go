package dto

// LOGIN DTO IS MODEL THAT USED BY CLIENT WHEN POST FROM /LOGIN URL
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
