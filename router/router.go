package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/rivendell/handlers"
	"github.com/jinzhu/gorm"
)

// Create a router engine
func Create(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := &handlers.Handler{Db: db}

	r.GET("/ping", h.HandlePing)
	r.GET("/database/ping", h.HandlePingDatabase)

	return r
}
