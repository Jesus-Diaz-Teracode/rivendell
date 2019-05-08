package models

// User is the main model
type User struct {
	ID        int64  `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique_index;not null"`
	Name      string
}
