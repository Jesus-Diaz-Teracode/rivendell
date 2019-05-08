package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/rivendell/handlers"
	"github.com/jinzhu/gorm"
)

// CreateEngine creates and configure the engine
func CreateEngine(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := &handlers.Handler{Db: db}

	r.GET("/ping", h.HandlePing)
	r.GET("/database/ping", h.HandlePingDatabase)

	r.POST("/users", h.CreateUser)
	r.GET("/users", h.ListUsers)

	return r
}
