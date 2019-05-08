package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlePing return an ok response
func (Handler) HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
