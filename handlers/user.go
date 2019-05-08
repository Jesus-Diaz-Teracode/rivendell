package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jedi4z/rivendell/models"
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

// ListUsers list all users from database
func (h *Handler) ListUsers(c *gin.Context) {
	var users []models.User
	results := h.Db.Find(&users)

	c.JSON(http.StatusOK, results)
}
