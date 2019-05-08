package handlers

import (
	"github.com/jinzhu/gorm"
)

// Handler is a struct to share
// dependencies with all handlers
type Handler struct {
	Db *gorm.DB
}
