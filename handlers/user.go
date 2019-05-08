package handlers

import (
	"net/http"

	"github.com/jedi4z/rivendell/models"

	"github.com/gin-gonic/gin"
)

// CreateUser handler the request to create a new user.
func (h *Handler) CreateUser(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Email:     json.Email,
	}

	h.Db.Create(&user)

	c.JSON(http.StatusOK, user)
}
