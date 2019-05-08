package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlePing return an ok response
func (Handler) HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// HandlePingDatabase return ok if the database is up
func (h *Handler) HandlePingDatabase(c *gin.Context) {
	err := h.Db.DB().Ping()
	if err != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{"message": "Error with the database connection"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
