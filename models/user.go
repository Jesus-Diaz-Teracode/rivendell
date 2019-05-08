package models

// User is the main model
type User struct {
	FirstName string `gorm:"not null" binding:"required" json:"first_name"`
	LastName  string `gorm:"not null" binding:"required" json:"last_name"`
	Email     string `gorm:"not null" binding:"required,email" json:"email"`
}
